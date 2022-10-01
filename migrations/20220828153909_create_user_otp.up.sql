CREATE TABLE IF NOT EXISTS user_otps
(
    id         serial PRIMARY KEY,
    user_id    INT                                 NOT NULL,
    otp        VARCHAR(10)                         NOT NULL,
    counter    INT                                 NOT NULL,
    mode       VARCHAR(100)                        NOT NULL,
    used       BOOLEAN                             NOT NULL,
    created_at TIMESTAMP default CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP default CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE SET NULL
);