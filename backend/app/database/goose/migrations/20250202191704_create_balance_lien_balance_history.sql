-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "balances" (
  "id" uuid constraint balances_pk primary key DEFAULT uuid_generate_v4(),
  "available_balance" bigint NOT NULL,
  "currency" varchar(50) NOT NULL,
  "change_amount" bigint NOT NULL,
  "locked_amount" bigint NOT NULL,
  "status" varchar(100) NOT NULL,
  "mode" varchar(100) NOT NULL,
  "hash" varchar(255) NOT NULL,
  "hash_key" varchar(255) NOT NULL,
  "previous_hash" varchar(255) DEFAULT null,
  "sequence" bigint NOT NULL,
  "transaction_id" uuid DEFAULT NULL,
  "tnx_time" timestamp default null,
  "created_at" timestamp default current_timestamp not null,
"updated_at" timestamp default null
);

CREATE TABLE IF NOT EXISTS  "balance_history" (
"id" uuid constraint balance_history_pk primary key DEFAULT uuid_generate_v4(),
 "balance_id" uuid NOT NULL,
  "available_balance" bigint NOT NULL,
  "currency" varchar(50) NOT NULL,
  "change_amount" bigint NOT NULL,
  "locked_amount" bigint NOT NULL,
  "status" varchar(100) NOT NULL,
  "mode" varchar(100) NOT NULL,
  "hash" varchar(255) NOT NULL,
  "hash_key" varchar(255) NOT NULL,   
  "previous_hash" varchar(255) DEFAULT null,
  "sequence" bigint NOT NULL,
  "transaction_id" uuid DEFAULT NULL,
  "operation" varchar(100) NOT NULL,
  "tnx_time" timestamp default null,
  "created_at" timestamp default current_timestamp not null,
"updated_at" timestamp default null
);

CREATE TABLE IF NOT EXISTS  "liens" (
"id" uuid constraint liens_pk primary key DEFAULT uuid_generate_v4(),
  "balance_id" uuid NOT NULL,
  "transaction_id" uuid NOT NULL,
  "lien_amount" bigint NOT NULL,
  "currency" varchar(50) NOT NULL,
  "status" varchar(100) NOT NULL,
  "created_at" timestamp DEFAULT current_timestamp not null,
  "updated_at" timestamp default null
);

ALTER TABLE "balance_history" ADD FOREIGN KEY ("balance_id") REFERENCES "balances" ("id");
ALTER TABLE "liens" ADD FOREIGN KEY ("balance_id") REFERENCES "balances" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table balances;
drop table balance_history;
drop table liens;
-- +goose StatementEnd
