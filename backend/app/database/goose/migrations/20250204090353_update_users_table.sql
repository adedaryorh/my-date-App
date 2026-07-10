-- +goose Up
-- +goose StatementBegin
ALTER TABLE "users" ADD COLUMN IF NOT EXISTS business_name varchar(256) DEFAULT NULL;
ALTER TABLE "users" ADD COLUMN IF NOT EXISTS industry_type varchar(256) DEFAULT NULL;
ALTER TABLE "users" ALTER COLUMN date_of_birth DROP NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE "users" DROP COLUMN business_name;
ALTER TABLE "users" DROP COLUMN industry_type;
ALTER TABLE "users" ALTER COLUMN date_of_birth SET NOT NULL;
-- +goose StatementEnd
