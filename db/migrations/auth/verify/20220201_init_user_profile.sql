-- Verify rasalibrary.auth:20220201_init_user_profile.sql on pg

BEGIN;

DO
$$
<<if_user_profile_table_exist_test>>
BEGIN
    IF NOT EXISTS(SELECT 1 FROM pg_tables WHERE schemaname = 'auth' AND tablename = 'user_profile') THEN
        RAISE EXCEPTION 'table auth.user_profile not found';
    END IF;
END if_user_profile_table_exist_test
$$;

ROLLBACK;
