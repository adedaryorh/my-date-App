-- +goose Up
-- +goose StatementBegin
Alter Table "users" ADD column next_login_at timestamp DEFAULT Null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Alter Table "users" DROP column next_login_at;
-- +goose StatementEnd
