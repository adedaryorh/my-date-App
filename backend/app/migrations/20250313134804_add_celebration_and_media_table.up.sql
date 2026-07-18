create table IF NOT EXISTS celebrations
(
	id uuid constraint celebrations_pk primary key DEFAULT uuid_generate_v4(),
	audience varchar(100) not null,
    celebration_type varchar(100) not null,
    owner varchar(100) not null,
	owner_identifier varchar(100) not null,
	owner_id varchar(100) not null,
    status varchar not null,
	frequency varchar(100) not null,
    created_by uuid not null,
    expires_at timestamp not null,
    longitude float(10) default null,
	latitude float(10) default null,
    notes text default null,
    selected_friends varchar[] default null,
	created_at timestamp default current_timestamp not null,
	updated_at timestamp default null
);

create table IF NOT EXISTS media
(
	id uuid constraint media_pk primary key DEFAULT uuid_generate_v4(),
	object_ref varchar(100) not null,
    object_id uuid not null,
    media_type varchar(100) not null,
    media_url varchar(100) not null,
	created_at timestamp default current_timestamp not null,
	updated_at timestamp default null
);

