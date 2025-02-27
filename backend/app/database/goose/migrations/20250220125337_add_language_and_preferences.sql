-- +goose Up
-- +goose StatementBegin
Alter Table "users" ADD column language varchar(100) DEFAULT Null;
Alter Table "users" ADD column notification_preference varchar[]DEFAULT Null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Alter Table "users" DROP column language;
Alter Table "users" DROP column notification_preference;
-- +goose StatementEnd
