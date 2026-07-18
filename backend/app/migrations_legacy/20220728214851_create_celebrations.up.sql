CREATE TABLE IF NOT EXISTS celebrations
(
    id             serial PRIMARY KEY,
    user_id        INT                                 NOT NULL,
    celebration_id VARCHAR(100)                        NOT NULL,
    message        VARCHAR(255),
    created_at     TIMESTAMP default CURRENT_TIMESTAMP NOT NULL,
    updated_at     TIMESTAMP default CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL
);

CREATE TABLE IF NOT EXISTS celebrations_media
(
    id             serial PRIMARY KEY,
    celebration_id INT                                 NOT NULL,
    type           VARCHAR(255),
    file_path      VARCHAR(255),
    file_extension VARCHAR(255),
    created_at     TIMESTAMP default CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_celebration FOREIGN KEY (celebration_id) REFERENCES celebrations (id) ON DELETE SET NULL
);