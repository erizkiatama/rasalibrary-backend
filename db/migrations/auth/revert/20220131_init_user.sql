-- Revert rasalibrary.auth:20223101_init_user from pg

BEGIN;

DROP TABLE IF EXISTS "auth".user;
DROP SCHEMA IF EXISTS "auth";

COMMIT;
