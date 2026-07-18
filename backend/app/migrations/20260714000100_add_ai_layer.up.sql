CREATE TABLE IF NOT EXISTS user_embeddings (
    user_id uuid PRIMARY KEY REFERENCES users(id) ON DELETE CASCADE,
    embedding double precision[] NOT NULL,
    profile_text text NOT NULL DEFAULT '',
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS moderation_queue (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    celebration_id uuid NOT NULL UNIQUE REFERENCES celebrations(id) ON DELETE CASCADE,
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    toxicity_score double precision NOT NULL CHECK (toxicity_score BETWEEN 0 AND 1),
    flagged_content text NOT NULL,
    status varchar(20) NOT NULL DEFAULT 'pending' CHECK (status IN ('pending', 'approved', 'rejected')),
    moderated_by uuid REFERENCES users(id) ON DELETE SET NULL,
    moderated_at timestamp,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS moderation_queue_status_created_idx
    ON moderation_queue(status, created_at DESC);

CREATE TABLE IF NOT EXISTS user_interactions (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id uuid NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    target_id uuid NOT NULL,
    target_type varchar(20) NOT NULL DEFAULT 'user' CHECK (target_type IN ('user', 'celebration')),
    action_type varchar(20) NOT NULL CHECK (action_type IN ('follow', 'skip', 'like', 'report')),
    metadata jsonb NOT NULL DEFAULT '{}'::jsonb,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS user_interactions_user_created_idx
    ON user_interactions(user_id, created_at DESC);
CREATE INDEX IF NOT EXISTS user_interactions_target_idx
    ON user_interactions(target_id, target_type);

ALTER TABLE users ADD COLUMN IF NOT EXISTS role varchar(20) NOT NULL DEFAULT 'user';

