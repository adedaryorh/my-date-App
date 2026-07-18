CREATE TABLE IF NOT EXISTS  "wallets" (
    "id" uuid constraint wallets_pk primary key DEFAULT uuid_generate_v4(),
    "balance_id" uuid NOT NULL,
    "currency" varchar(50) NOT NULL,
    "owner_type" varchar(50) NOT NULL,
    "owner_id" varchar(50) NOT NULL,
    "status" varchar(100) NOT NULL,
    "created_at" timestamp DEFAULT current_timestamp not null,
    "updated_at" timestamp default null
);

