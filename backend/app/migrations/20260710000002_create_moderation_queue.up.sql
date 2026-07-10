-- Create moderation_queue table for flagged celebrations awaiting review
CREATE TABLE IF NOT EXISTS moderation_queue
(
    id              SERIAL PRIMARY KEY,
    post_id         INTEGER                                 NOT NULL,
    user_id         INTEGER                                 NOT NULL,
    toxicity_score  REAL                                    NOT NULL,
    flagged_content TEXT,  -- The content that was flagged
    status          VARCHAR(50) DEFAULT 'pending' NOT NULL,  -- pending, approved, rejected
    moderated_by    INTEGER,  -- User ID of moderator who reviewed it
    moderated_at    TIMESTAMP NULL,
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_post FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE CASCADE,
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_moderator FOREIGN KEY (moderated_by) REFERENCES users (id) ON DELETE SET NULL
);

-- Create indexes for common queries
CREATE INDEX IF NOT EXISTS idx_moderation_queue_status ON moderation_queue(status);
CREATE INDEX IF NOT EXISTS idx_moderation_queue_created_at ON moderation_queue(created_at);
CREATE INDEX IF NOT EXISTS idx_moderation_queue_user_id ON moderation_queue(user_id);