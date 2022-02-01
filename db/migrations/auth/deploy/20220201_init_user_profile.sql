-- Deploy rasalibrary.auth:20220201_init_user_profile.sql to pg

BEGIN;

CREATE TABLE IF NOT EXISTS "auth".user_profile(
    id              SERIAL          PRIMARY KEY,
    name            VARCHAR(100)    NOT NULL,
    dob             DATE            NOT NULL,
    address         TEXT            NOT NULL,
    sex             VARCHAR(10)     NOT NULL,
    phone_number    VARCHAR(15)     NOT NULL,
    profile_photo   TEXT,
    user_id         INT             NOT NULL,
    created_at      TIMESTAMPTZ     DEFAULT now(),
    updated_at      TIMESTAMPTZ     DEFAULT now()
);

COMMENT ON TABLE "auth".user_profile IS 'User profile detail information';

ALTER TABLE "auth".user_profile ADD CONSTRAINT fk_user FOREIGN KEY (user_id) 
REFERENCES "auth".user(id) ON DELETE CASCADE ON UPDATE CASCADE;

ALTER TABLE "auth".user_profile ADD CONSTRAINT user_profile_user_id_key UNIQUE (user_id);

COMMIT;
