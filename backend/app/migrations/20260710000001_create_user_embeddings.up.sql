-- Create user_embeddings table for storing user profile embeddings
CREATE TABLE IF NOT EXISTS user_embeddings
(
    id          SERIAL PRIMARY KEY,
    user_id     INTEGER                                 NOT NULL,
    embedding   VECTOR(384),  -- Dimension for all-MiniLM-L6-v2 model
    bio_text    TEXT,
    interests   TEXT[],  -- PostgreSQL array type for interests
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

-- Create index on user_id for faster lookups
CREATE INDEX IF NOT EXISTS idx_user_embeddings_user_id ON user_embeddings(user_id);