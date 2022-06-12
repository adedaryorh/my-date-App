CREATE TABLE IF NOT EXISTS client_tokens
(
    id         serial PRIMARY KEY,
    token      VARCHAR(255)                        NOT NULL,
    expiry     TIMESTAMP                           NOT NULL,
    created_at TIMESTAMP default CURRENT_TIMESTAMP NOT NULL
);