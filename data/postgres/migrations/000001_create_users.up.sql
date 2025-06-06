CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS user (
    id bigserial PRIMARY KEY,
    email varchar(255) UNIQUE NOT NULL,
    username varchar(255) NOT NULL,
    password bytea NOT NULL,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW()
    );