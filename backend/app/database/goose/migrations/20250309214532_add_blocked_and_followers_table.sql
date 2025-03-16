-- +goose Up
-- +goose StatementBegin
Alter Table "users" ADD column push_notification_settings jsonb DEFAULT Null;
Alter Table "users" ADD column content_settings jsonb DEFAULT Null;
Alter Table "users" ADD column banned_words varchar[] DEFAULT Null;

create table IF NOT EXISTS followers
(
	id uuid constraint followers_pk primary key DEFAULT uuid_generate_v4(),
	user_id uuid not null,
    follower_id uuid not null,
  	created_at timestamp default current_timestamp not null,
	updated_at timestamp default null
);

create table IF NOT EXISTS blocked
(
	id uuid constraint blocked_pk primary key DEFAULT uuid_generate_v4(),
	user_id uuid not null,
    blocked_user_id uuid not null,
  	created_at timestamp default current_timestamp not null,
	updated_at timestamp default null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Alter Table "users" DROP column push_notification_settings;
Alter Table "users" DROP column content_settings;
Alter Table "users" DROP column banned_words;

drop table followers;
drop table blocked;
-- +goose StatementEnd
