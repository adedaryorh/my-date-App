-- +goose Up
-- +goose StatementBegin
Alter Table "users" ADD column address text DEFAULT Null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Alter Table "users" DROP column address;
-- +goose StatementEnd
