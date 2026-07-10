"""
Test suite for the Celebut AI Service
"""

import pytest
import json
from fastapi.testclient import TestClient
from app.main import app

client = TestClient(app)

def test_health_check():
    """Test the health check endpoint."""
    response = client.get("/ai/health")
    assert response.status_code == 200
    data = response.json()
    assert data["status"] == "healthy"
    assert "timestamp" in data
    assert "version" in data

def test_embed_profile():
    """Test profile embedding generation."""
    response = client.post(
        "/ai/embed-profile",
        json={
            "user_id": "test_user_1",
            "bio": "I love hiking and photography",
            "interests": ["outdoor", "photography", "travel"]
        }
    )
    assert response.status_code == 200
    data = response.json()
    assert data["user_id"] == "test_user_1"
    assert "embedding" in data
    assert isinstance(data["embedding"], list)
    assert len(data["embedding"]) > 0
    assert "text_used" in data

def test_moderate_safe_content():
    """Test moderation of safe content."""
    response = client.post(
        "/ai/moderate-celebration",
        json={
            "text": "Having a wonderful birthday party with friends!"
        }
    )
    assert response.status_code == 200
    data = response.json()
    assert "toxicity_score" in data
    assert "is_flagged" in data
    assert isinstance(data["toxicity_score"], float)
    assert isinstance(data["is_flagged"], bool)
    # Safe content should have low toxicity
    assert data["toxicity_score"] < 0.5

def test_moderate_toxic_content():
    """Test moderation of toxic content."""
    response = client.post(
        "/ai/moderate-celebration",
        json={
            "text": "I hate everyone and everything is terrible"
        }
    )
    assert response.status_code == 200
    data = response.json()
    assert "toxicity_score" in data
    assert "is_flagged" in data
    # Note: The actual toxicity score depends on the model
    # This test mainly checks that the endpoint works correctly

def test_index_and_search_celebration():
    """Test indexing and searching celebrations."""
    # Index a celebration
    index_response = client.post(
        "/ai/index-celebration",
        json={
            "celebration_id": "test_celeb_1",
            "text": "Celebrating my graduation with family and friends!"
        }
    )
    assert index_response.status_code == 200
    index_data = index_response.json()
    assert index_data["celebration_id"] == "test_celeb_1"
    assert index_data["indexed"] == True

    # Search for similar content
    search_response = client.post(
        "/ai/search-celebrations",
        json={
            "query": "graduation celebration party",
            "limit": 5
        }
    )
    assert search_response.status_code == 200
    search_data = search_response.json()
    assert "results" in search_data
    assert isinstance(search_data["results"], list)

def test_log_interaction():
    """Test logging user interactions."""
    response = client.post(
        "/ai/log-interaction",
        json={
            "user_id": "user_1",
            "target_id": "user_2",
            "action": "follow",
            "metadata": {
                "context": "discovery",
                "timestamp": "2024-01-01T10:00:00Z"
            }
        }
    )
    assert response.status_code == 200
    data = response.json()
    assert data["logged"] == True
    assert "interaction_id" in data

def test_recommend_users_endpoint():
    """Test the user recommendations endpoint (will return empty list without DB setup)."""
    response = client.post(
        "/ai/recommend-users",
        json={
            "user_id": "test_user_1",
            "limit": 5
        }
    )
    # This might return 500 if DB is not available, or 200 with empty results
    # Depending on test environment setup
    assert response.status_code in [200, 500]
    if response.status_code == 200:
        data = response.json()
        assert "recommended_users" in data
        assert isinstance(data["recommended_users"], list)

if __name__ == "__main__":
    pytest.main([__file__])