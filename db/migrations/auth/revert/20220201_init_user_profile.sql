-- Revert rasalibrary.auth:20220201_init_user_profile.sql from pg

BEGIN;

DROP TABLE IF EXISTS "auth".user_profile;

COMMIT;
