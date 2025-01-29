-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table IF NOT EXISTS users
(
	id uuid constraint users_pk primary key DEFAULT uuid_generate_v4(),
	first_name varchar(100) not null,
    last_name varchar(100) not null,
    middle_name varchar(100),
	email varchar(100) not null UNIQUE,
    status varchar not null,
	password_hash text not null,
    kyc jsonb not null,
	created_at timestamp default current_timestamp not null,
	updated_at timestamp default null
);

create unique index users_email_uindex on users (email);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd