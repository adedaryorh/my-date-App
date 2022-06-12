CREATE TABLE IF NOT EXISTS industries
(
    id          serial PRIMARY KEY,
    name        VARCHAR(255)                        NOT NULL,
    description TEXT                                NOT NULL,
    created_at  TIMESTAMP default CURRENT_TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP default CURRENT_TIMESTAMP NOT NULL
);