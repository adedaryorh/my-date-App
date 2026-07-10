"""
Celebut AI Service
A microservice for AI-powered features including user recommendations,
content moderation, semantic search, and feedback learning.
"""

import os
import logging
from contextlib import asynccontextmanager
from typing import List, Optional, Dict, Any
import numpy as np

from fastapi import FastAPI, HTTPException, Depends, BackgroundTasks
from fastapi.middleware.cors import CORSMiddleware
from pydantic import BaseModel, Field
import torch
from sentence_transformers import SentenceTransformer
import faiss
import numpy as np
from transformers import pipeline
from sklearn.metrics.pairwise import cosine_similarity
import psycopg2
from psycopg2.extras import RealDictCursor
import hashlib
import json
from datetime import datetime

# Load environment variables
from dotenv import load_dotenv
load_dotenv()

# Configure logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

# Configuration
DATABASE_URL = os.getenv("DATABASE_URL", "postgresql://localhost/celebut")
MODEL_CACHE_DIR = os.getenv("MODEL_CACHE_DIR", "./models")
EMBEDDING_MODEL_NAME = os.getenv("EMBEDDING_MODEL_NAME", "all-MiniLM-L6-v2")
TOXICITY_MODEL_NAME = os.getenv("TOXICITY_MODEL_NAME", "unitary/toxic-bert")
FAISS_INDEX_PATH = os.getenv("FAISS_INDEX_PATH", "./faiss_index")
TOXICITY_THRESHOLD = float(os.getenv("TOXICITY_THRESHOLD", "0.8"))

# Global variables for models and indexes
embedding_model: Optional[SentenceTransformer] = None
toxicity_classifier: Optional[Any] = None
celebration_index: Optional[faiss.Index] = None
celebration_id_map: Dict[int, str] = {}  # Maps FAISS index to celebration ID
user_embeddings_cache: Dict[str, np.ndarray] = {}  # Cache for user embeddings

@asynccontextmanager
async def lifespan(app: FastAPI):
    """Initialize models and load existing data on startup."""
    global embedding_model, toxicity_classifier, celebration_index

    logger.info("Starting up AI Service...")

    # Initialize embedding model
    logger.info(f"Loading embedding model: {EMBEDDING_MODEL_NAME}")
    embedding_model = SentenceTransformer(EMBEDDING_MODEL_NAME, cache_folder=MODEL_CACHE_DIR)

    # Initialize toxicity classifier
    logger.info(f"Loading toxicity model: {TOXICITY_MODEL_NAME}")
    toxicity_classifier = pipeline(
        "text-classification",
        model=TOXICITY_MODEL_NAME,
        tokenizer=TOXICITY_MODEL_NAME,
        return_all_scores=True
    )

    # Initialize FAISS index for celebration embeddings
    dimension = embedding_model.get_sentence_embedding_dimension()
    celebration_index = faiss.IndexFlatIP(dimension)  # Inner product for cosine similarity

    # Load existing celebration embeddings from database if available
    await load_celebration_embeddings()

    logger.info("AI Service startup complete")

    yield

    logger.info("Shutting down AI Service...")
    # Save FAISS index on shutdown
    if celebration_index is not None:
        faiss.write_index(celebration_index, FAISS_INDEX_PATH)
        logger.info(f"Saved FAISS index to {FAISS_INDEX_PATH}")

app = FastAPI(
    title="Celebut AI Service",
    description="AI-powered features for Celebut social platform",
    version="1.0.0",
    lifespan=lifespan
)

# Add CORS middleware
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],  # Configure appropriately for production
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

# Pydantic models for request/response validation
class ProfileEmbeddingRequest(BaseModel):
    user_id: str
    bio: Optional[str] = ""
    interests: List[str] = []

class ProfileEmbeddingResponse(BaseModel):
    user_id: str
    embedding: List[float]
    text_used: str

class UserRecommendationRequest(BaseModel):
    user_id: str
    limit: int = 10
    exclude_followed: bool = True
    exclude_blocked: bool = True

class UserRecommendationResponse(BaseModel):
    recommended_users: List[Dict[str, Any]]

class ModerationRequest(BaseModel):
    text: str

class ModerationResponse(BaseModel):
    toxicity_score: float
    is_flagged: bool
    categories: dict

class CelebrationIndexRequest(BaseModel):
    celebration_id: str
    text: str

class CelebrationIndexResponse(BaseModel):
    celebration_id: str
    indexed: bool

class SearchRequest(BaseModel):
    query: str
    limit: int = 10

class SearchResponse(BaseModel):
    results: List[Dict[str, Any]]

class InteractionLogRequest(BaseModel):
    user_id: str
    target_id: str
    action: str  # follow, skip, like, report
    metadata: Optional[Dict[str, Any]] = None

class InteractionLogResponse(BaseModel):
    logged: bool
    interaction_id: Optional[str] = None

class HealthResponse(BaseModel):
    status: str
    timestamp: str
    version: str

# Dependency to get database connection
def get_db_connection():
    """Get a database connection."""
    try:
        conn = psycopg2.connect(DATABASE_URL)
        return conn
    except Exception as e:
        logger.error(f"Database connection error: {e}")
        raise HTTPException(status_code=500, detail="Database connection failed")

# Helper functions
def get_text_embedding(text: str) -> np.ndarray:
    """Generate embedding for text using the sentence transformer model."""
    if not text.strip():
        # Return zero vector for empty text
        return np.zeros(embedding_model.get_sentence_embedding_dimension())

    embedding = embedding_model.encode(text, normalize_embeddings=True)
    return embedding

def get_profile_text(bio: str, interests: List[str]) -> str:
    """Combine bio and interests into a single text for embedding."""
    parts = []
    if bio.strip():
        parts.append(bio.strip())
    if interests:
        interests_text = ", ".join(interests)
        parts.append(f"Interests: {interests_text}")
    return " ".join(parts)

def get_toxicity_score(text: str) -> tuple[float, dict]:
    """Get toxicity score and categories from the toxicity model."""
    if not text.strip():
        return 0.0, {}

    results = toxicity_classifier(text)
    # The toxic-bert model returns labels like 'toxic', 'severe_toxic', etc.
    toxicity_scores = {}
    max_score = 0.0

    for result in results[0]:  # results is a list of lists
        label = result['label']
        score = result['score']
        toxicity_scores[label] = score
        if score > max_score:
            max_score = score

    return max_score, toxicity_scores

async def load_celebration_embeddings():
    """Load existing celebration embeddings from database into FAISS index."""
    try:
        conn = get_db_connection()
        cursor = conn.cursor(cursor_factory=RealDictCursor)

        # Try to get celebrations with their text content
        # Adjust table/column names based on actual schema
        cursor.execute("""
            SELECT id, COALESCE(message, '') as message
            FROM posts
            WHERE message IS NOT NULL AND message != ''
        """)

        celebrations = cursor.fetchall()

        if celebrations:
            texts = [celebr['message'] for celeb in celebrations]
            ids = [str(celeb['id']) for celeb in celebrations]

            # Generate embeddings
            embeddings = embedding_model.encode(texts, normalize_embeddings=True)

            # Add to FAISS index
            celebration_index.add(embeddings.astype('float32'))

            # Create mapping from FAISS index to celebration ID
            for i, celeb_id in enumerate(ids):
                celebration_id_map[i] = celeb_id

            logger.info(f"Loaded {len(celebrations)} celebration embeddings into FAISS index")

        cursor.close()
        conn.close()
    except Exception as e:
        logger.warning(f"Could not load existing celebration embeddings: {e}")
        # Continue with empty index

# API Endpoints

@app.get("/ai/health", response_model=HealthResponse)
async def health_check():
    """Health check endpoint."""
    return HealthResponse(
        status="healthy",
        timestamp=datetime.utcnow().isoformat(),
        version="1.0.0"
    )

@app.post("/ai/embed-profile", response_model=ProfileEmbeddingResponse)
async def embed_profile(request: ProfileEmbeddingRequest):
    """Generate and store embedding for a user's profile."""
    try:
        # Combine bio and interests for embedding
        profile_text = get_profile_text(request.bio, request.interests)

        # Generate embedding
        embedding = get_text_embedding(profile_text)

        # Cache the embedding
        user_embeddings_cache[request.user_id] = embedding

        # Optionally store in database for persistence
        # This would require a user_embeddings table

        return ProfileEmbeddingResponse(
            user_id=request.user_id,
            embedding=embedding.tolist(),
            text_used=profile_text
        )
    except Exception as e:
        logger.error(f"Error generating profile embedding: {e}")
        raise HTTPException(status_code=500, detail=str(e))

@app.post("/ai/recommend-users", response_model=UserRecommendationResponse)
async def recommend_users(request: UserRecommendationRequest):
    """Recommend users to follow based on profile similarity, interests, and interaction history."""
    try:
        # Get the source user's embedding
        if request.user_id not in user_embeddings_cache:
            # Try to generate it from database if not in cache
            conn = get_db_connection()
            cursor = conn.cursor(cursor_factory=RealDictCursor)
            cursor.execute(
                """
                SELECT bio, interests FROM users WHERE id = %s
                """,
                (request.user_id,)
            )
            user_data = cursor.fetchone()
            cursor.close()
            conn.close()

            if not user_data:
                raise HTTPException(status_code=404, detail="User not found")

            bio = user_data.get('bio') or ''
            interests = user_data.get('interests') or []
            profile_text = get_profile_text(bio, interests)
            embedding = get_text_embedding(profile_text)
            user_embeddings_cache[request.user_id] = embedding
        else:
            embedding = user_embeddings_cache[request.user_id]

        # Get all other users from database
        conn = get_db_connection()
        cursor = conn.cursor(cursor_factory=RealDictCursor)

        # Exclude current user and potentially followed/blocked users
        query = """
            SELECT id, bio, interests, username, first_name, last_name, profile_image_url
            FROM users
            WHERE id != %s
        """
        params = [request.user_id]

        if request.exclude_followed:
            # Exclude users that the current user already follows
            query += """
                AND id NOT IN (
                    SELECT followed_id FROM followers WHERE follower_id = %s
                )
            """
            params.append(request.user_id)

        if request.exclude_blocked:
            # Exclude users that the current user has blocked
            query += """
                AND id NOT IN (
                    SELECT blocked_id FROM blocks WHERE blocker_id = %s
                )
            """
            params.append(request.user_id)

        cursor.execute(query, params)
        other_users = cursor.fetchall()
        cursor.close()
        conn.close()

        if not other_users:
            return UserRecommendationResponse(recommended_users=[])

        # Calculate similarities
        user_texts = []
        for user in other_users:
            user_text = get_profile_text(
                user.get('bio') or '',
                user.get('interests') or []
            )
            user_texts.append(user_text)

        if user_texts:
            # Batch encode for efficiency
            other_embeddings = embedding_model.encode(user_texts, normalize_embeddings=True)

            # Calculate cosine similarity
            similarities = cosine_similarity([embedding], other_embeddings)[0]

            # Boost score based on shared interests
            # Get current user's interests for bonus
            conn = get_db_connection()
            cursor = conn.cursor(cursor_factory=RealDictCursor)
            cursor.execute(
                """
                SELECT interests FROM users WHERE id = %s
                """,
                (request.user_id,)
            )
            curr_user = cursor.fetchone()
            cursor.close()
            conn.close()
            curr_interests = set(curr_user.get('interests') or []) if curr_user else set()

            for i, user in enumerate(other_users):
                user_interests = set(user.get('interests') or [])
                shared = len(curr_interests.intersection(user_interests))
                interest_boost = 0.05 * shared  # small boost per shared interest
                similarities[i] += interest_boost

        # Adjust scores based on interaction history (follow, like, skip, report)
        # Get interactions where current user is the actor
        if other_users:
            target_user_ids = [str(user['id']) for user in other_users]
            # Fetch interactions
            conn = get_db_connection()
            cursor = conn.cursor(cursor_factory=RealDictCursor)
            # We'll fetch recent interactions; limit to last 100 per action type maybe
            cursor.execute(
                """
                SELECT target_id, action_type, COUNT(*) as count
                FROM user_interactions
                WHERE user_id = %s AND action_type IN ('follow', 'like', 'skip', 'report')
                GROUP BY target_id, action_type
                """,
                (request.user_id,)
            )
            interaction_rows = cursor.fetchall()
            cursor.close()
            conn.close()

            # Compute adjustment per target
            adjustment_dict = {}
            # Define weights
            weights = {
                'follow': 0.2,
                'like': 0.1,
                'skip': -0.15,
                'report': -0.25
            }
            for row in interaction_rows:
                target_id = row['target_id']
                action = row['action_type']
                count = row['count']
                # Cap count to avoid excessive boost
                weight = weights.get(action, 0.0) * min(count, 5)  # cap at 5 occurrences
                adjustment_dict[target_id] = adjustment_dict.get(target_id, 0.0) + weight

            # Apply adjustments
            for i, user in enumerate(other_users):
                target_id = str(user['id'])
                adjustment = adjustment_dict.get(target_id, 0.0)
                similarities[i] += adjustment
                # Optional: clip similarity to [0, 1] if desired
                # similarities[i] = max(0.0, min(1.0, similarities[i]))

        # Sort by similarity score (descending)
        scored_users = list(zip(other_users, similarities))
        scored_users.sort(key=lambda x: x[1], reverse=True)

        # Take top results
        top_users = scored_users[:request.limit]

        # Format response
        recommended_users = []
        for user, score in top_users:
            recommended_users.append({
                "user_id": str(user['id']),
                "username": user['username'],
                "first_name": user['first_name'],
                "last_name": user['last_name'],
                "profile_image_url": user['profile_image_url'],
                "similarity_score": float(score),
                "bio": user['bio'],
                "interests": user['interests']
            })

        return UserRecommendationResponse(recommended_users=recommended_users)

    except Exception as e:
        logger.error(f"Error generating user recommendations: {e}")
        raise HTTPException(status_code=500, detail=str(e))

@app.post("/ai/moderate-celebration", response_model=ModerationResponse)
async def moderate_celebration(request: ModerationRequest):
    """Check celebration content for toxicity/inappropriate content."""
    try:
        toxicity_score, categories = get_toxicity_score(request.text)
        is_flagged = toxicity_score >= TOXICITY_THRESHOLD

        return ModerationResponse(
            toxicity_score=toxicity_score,
            is_flagged=is_flagged,
            categories=categories
        )
    except Exception as e:
        logger.error(f"Error moderating content: {e}")
        raise HTTPException(status_code=500, detail=str(e))

@app.post("/ai/index-celebration", response_model=CelebrationIndexResponse)
async def index_celebration(request: CelebrationIndexRequest):
    """Index a celebration's text for semantic search."""
    try:
        if not request.text.strip():
            raise HTTPException(status_code=400, detail="Text cannot be empty")

        # Generate embedding
        embedding = get_text_embedding(request.text)

        # Add to FAISS index
        celebration_index.add(embedding.reshape(1, -1).astype('float32'))

        # Get the index position where this was added
        index_position = celebration_index.ntotal - 1
        celebration_id_map[index_position] = request.celebration_id

        logger.info(f"Indexed celebration {request.celebration_id} at position {index_position}")

        return CelebrationIndexResponse(
            celebration_id=request.celebration_id,
            indexed=True
        )
    except Exception as e:
        logger.error(f"Error indexing celebration: {e}")
        raise HTTPException(status_code=500, detail=str(e))

@app.post("/ai/search-celebrations", response_model=SearchResponse)
async def search_celebrations(request: SearchRequest):
    """Search for celebrations using semantic similarity."""
    try:
        if not request.query.strip():
            raise HTTPException(status_code=400, detail="Query cannot be empty")

        if celebration_index.ntotal == 0:
            return SearchResponse(results=[])

        # Generate query embedding
        query_embedding = get_text_embedding(request.query)

        # Search in FAISS index
        k = min(request.limit, celebration_index.ntotal)
        scores, indices = celebrity_index.search(
            query_embedding.reshape(1, -1).astype('float32'), k
        )

        # Get celebration IDs from indices
        celebration_ids = []
        valid_indices = []

        for i, idx in enumerate(indices[0]):
            if idx != -1 and idx in celebration_id_map:  # FAISS uses -1 for empty slots
                celebration_ids.append(celebration_id_map[int(idx)])
                valid_indices.append(i)

        if not celebration_ids:
            return SearchResponse(results=[])

        # Fetch celebration details from database
        conn = get_db_connection()
        cursor = conn.cursor(cursor_factory=RealDictCursor)

        # Using ANY for efficient querying
        cursor.execute(
            """
            SELECT id, message, user_id, created_at,
                   (SELECT json_build_object(
                        'id', u.id,
                        'username', u.username,
                        'first_name', u.first_name,
                        'last_name', u.last_name
                       ) as user
            FROM posts p
            JOIN users u ON p.user_id = u.id
            WHERE id = ANY(%s)
            ORDER BY created_at DESC
            """,
            (celebration_ids,)
        )

        celebrations = cursor.fetchall()
        cursor.close()
        conn.close()

        # Format results with similarity scores
        results = []
        score_idx = 0

        for celeb in celebrations:
            # Find the corresponding score
            while score_idx < len(indices[0]) and int(indices[0][score_idx]) not in celebration_id_map:
                score_idx += 1

            if score_idx < len(indices[0]):
                similarity_score = float(scores[0][score_idx])
                score_idx += 1
            else:
                similarity_score = 0.0

            results.append({
                "celebration_id": str(celeb['id']),
                "message": celeb['message'],
                "user_id": str(celeb['user_id']),
                "created_at": celeb['created_at'].isoformat() if celeb['created_at'] else None,
                "user": celeb['user'],
                "similarity_score": similarity_score
            })

        return SearchResponse(results=results)
    except Exception as e:
        logger.error(f"Error searching celebrations: {e}")
        raise HTTPException(status_code=500, detail=str(e))

@app.post("/ai/log-interaction", response_model=InteractionLogResponse)
async def log_interaction(request: InteractionLogRequest):
    """Log user interaction for feedback loop."""
    try:
        # Validate action type
        valid_actions = ["follow", "skip", "like", "report"]
        if request.action not in valid_actions:
            raise HTTPException(
                status_code=400,
                detail=f"Invalid action. Must be one of {valid_actions}"
            )

        # Store in database
        conn = get_db_connection()
        cursor = conn.cursor()

        cursor.execute(
            """
            INSERT INTO user_interactions
            (user_id, target_id, action_type, metadata, created_at)
            VALUES (%s, %s, %s, %s, %s)
            RETURNING id
            """,
            (
                request.user_id,
                request.target_id,
                request.action,
                json.dumps(request.metadata or {}),
                datetime.utcnow()
            )
        )

        interaction_id = cursor.fetchone()[0]
        conn.commit()
        cursor.close()
        conn.close()

        logger.info(f"Logged interaction: {request.user_id} {request.action} {request.target_id}")

        return InteractionLogResponse(
            logged=True,
            interaction_id=str(interaction_id)
        )
    except Exception as e:
        logger.error(f"Error logging interaction: {e}")
        raise HTTPException(status_code=500, detail=str(e))

if __name__ == "__main__":
    import uvicorn
    uvicorn.run(app, host="0.0.0.0", port=8000)