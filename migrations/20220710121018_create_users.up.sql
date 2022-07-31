CREATE TABLE IF NOT EXISTS users
(
    id                  serial PRIMARY KEY,
    user_id             VARCHAR(100)                        NOT NULL,
    first_name          VARCHAR(255),
    last_name           VARCHAR(255),
    username            VARCHAR(255),
    business_name       VARCHAR(255),
    email               VARCHAR(255)                        NOT NULL,
    country_code        VARCHAR(5)                          NOT NULL,
    phone               VARCHAR(20)                         NOT NULL,
    gender              VARCHAR(20),
    dob                 DATE,
    relationship_status VARCHAR(50),
    interests           VARCHAR(255),
    industry_id         int,
    account_type_id     int,
    password_hash       VARCHAR(255)                        NOT NULL,
    status              VARCHAR(100)                        NOT NULL,
    created_at          TIMESTAMP default CURRENT_TIMESTAMP NOT NULL,
    updated_at          TIMESTAMP default CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_industry FOREIGN KEY (industry_id) REFERENCES industries (id) ON DELETE SET NULL,
    CONSTRAINT fk_account_type FOREIGN KEY (account_type_id) REFERENCES account_types (id) ON DELETE SET DEFAULT
);
