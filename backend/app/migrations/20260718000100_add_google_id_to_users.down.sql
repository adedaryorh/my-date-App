DROP INDEX IF EXISTS users_google_id_uindex;
ALTER TABLE users DROP COLUMN IF EXISTS google_id;
