CREATE TABLE IF NOT EXISTS user_reactions
(
    id         serial PRIMARY KEY,
    user_id    INT                                 NOT NULL,
    post_id    INT                                 NOT NULL,
    reaction   VARCHAR(100)                        NOT NULL,
    created_at TIMESTAMP default CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL,
    CONSTRAINT fk_post FOREIGN KEY (post_id) REFERENCES posts (id) ON DELETE SET NULL
);
