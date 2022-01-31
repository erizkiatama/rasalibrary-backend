-- Deploy rasalibrary.auth:20223101_init_user to pg

BEGIN;

CREATE SCHEMA IF NOT EXISTS "auth";

CREATE TABLE IF NOT EXISTS "auth".user(
    id          SERIAL          PRIMARY KEY,
    email       VARCHAR(100)    NOT NULL UNIQUE,
    password    VARCHAR(100)    NOT NULL,
    is_admin    BOOLEAN         NOT NULL,
    created_at  TIMESTAMPTZ     DEFAULT now(),
    updated_at  TIMESTAMPTZ     DEFAULT now()
);

COMMENT ON TABLE "auth".user IS 'User account basic information';

COMMIT;
