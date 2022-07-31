CREATE TABLE IF NOT EXISTS account_types
(
    id          serial PRIMARY KEY,
    name        VARCHAR(100),
    description VARCHAR(255),
    created_at  TIMESTAMP default CURRENT_TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP default CURRENT_TIMESTAMP NOT NULL
);

-- IMPORTANT: Seed account types during migration
INSERT INTO account_types(name, description)
VALUES ('Basic', 'A basic user/an individual'),
       ('Business', 'A business account'),
       ('Admin', 'An administrative account');
