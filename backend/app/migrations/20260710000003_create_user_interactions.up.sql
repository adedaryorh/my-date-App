-- Create user_interactions table for tracking user actions (follow, skip, like, report)
CREATE TABLE IF NOT EXISTS user_interactions
(
    id              SERIAL PRIMARY KEY,
    user_id         INTEGER                                 NOT NULL,  -- The user performing the action
    target_id       INTEGER                                 NOT NULL,  -- The user or post being acted upon
    target_type     VARCHAR(20)  DEFAULT 'user' NOT NULL,  -- 'user' or 'post'
    action_type     VARCHAR(20)  NOT NULL,  -- follow, skip, like, report
    metadata        JSONB,  -- Additional context about the interaction
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CONSTRAINT fk_target_user FOREIGN KEY (target_id) REFERENCES users (id) ON DELETE CASCADE,
    -- Note: We don't add a foreign key to posts here because target_type can be 'user' or 'post'
    -- Application logic will handle validating the target_id based on target_type
);

-- Create indexes for common queries
CREATE INDEX IF NOT EXISTS idx_user_interactions_user_id ON user_interactions(user_id);
CREATE INDEX IF NOT EXISTS idx_user_interactions_target_id ON user_interactions(target_id);
CREATE INDEX IF NOT EXISTS idx_user_interactions_action_type ON user_interactions(action_type);
CREATE INDEX IF NOT EXISTS idx_user_interactions_created_at ON user_interactions(created_at);
CREATE INDEX IF NOT EXISTS idx_user_interactions_target_type ON user_interactions(target_type);