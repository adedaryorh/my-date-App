CREATE TABLE IF NOT EXISTS user_relationships
(
    id               serial PRIMARY KEY,
    sender_user_id   INT                                 NOT NULL,
    receiver_user_id INT                                 NOT NULL,
    status           VARCHAR(100)                        NOT NULL,
    created_at       TIMESTAMP default CURRENT_TIMESTAMP NOT NULL,
    updated_at       TIMESTAMP default CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_sender_user FOREIGN KEY (sender_user_id) REFERENCES users (id) ON DELETE SET NULL,
    CONSTRAINT fk_receiver_user FOREIGN KEY (receiver_user_id) REFERENCES users (id) ON DELETE SET NULL
);
