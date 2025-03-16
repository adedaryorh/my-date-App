-- +goose Up
-- +goose StatementBegin
Alter Table "celebrations" ADD column caption text not Null;
Alter Table "celebrations" ADD column celebration_kind varchar(256) not Null;
Alter Table "celebrations" ADD column celebration_date timestamp not Null;

Alter Table "media" ADD column owner_id varchar(256) not Null;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Alter Table "celebrations" DROP column caption;
Alter Table "celebrations" DROP column celebration_kind;
Alter Table "celebrations" DROP column celebration_date;

Alter Table "media" DROP column owner_id;
-- +goose StatementEnd
