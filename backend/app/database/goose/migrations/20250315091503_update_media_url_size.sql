-- +goose Up
-- +goose StatementBegin
Alter Table "media" DROP column media_url;
Alter Table "media" ADD column media_url varchar(256) not Null;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Alter Table "media" DROP column media_url;
Alter Table "media" ADD column media_url varchar(100) not Null;
-- +goose StatementEnd
