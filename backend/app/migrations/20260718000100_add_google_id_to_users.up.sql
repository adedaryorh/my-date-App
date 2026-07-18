ALTER TABLE users
    ADD COLUMN IF NOT EXISTS google_id varchar(255) NOT NULL DEFAULT '';

CREATE UNIQUE INDEX IF NOT EXISTS users_google_id_uindex
    ON users (google_id)
    WHERE google_id <> '';

