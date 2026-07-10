-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table IF NOT EXISTS users
(
	id uuid constraint users_pk primary key DEFAULT uuid_generate_v4(),
	first_name varchar(100) not null,
    last_name varchar(100) default null,
    username varchar(100)not null UNIQUE,
	email varchar(100) not null UNIQUE,
	country_code varchar(100) not null,
	longitude float(10) default null,
	latitude float(10) default null,
	phone_number varchar(100) not null,
	completion_state int not null,
	ip_address varchar(100) default null,
	device_type varchar(100) default null,
	date_of_birth timestamp not Null,
	account_type varchar(100) not null,
	interests varchar[] DEFAULT null,
	profile_image_url varchar(256) default null,
	verification_status varchar not null,
    status varchar not null,
	password_hash text not null,
	created_at timestamp default current_timestamp not null,
	updated_at timestamp default null
);

create unique index users_email_uindex on users (email);
create unique index users_username_uindex on users (username);
create unique index users_phone_number_uindex on users (phone_number);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd