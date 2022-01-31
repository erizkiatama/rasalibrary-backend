-- Verify rasalibrary.auth:20223101_init_user on pg

BEGIN;

DO
$$
    <<if_auth_schema_exist_test>>
    BEGIN
        IF NOT EXISTS(SELECT 1 FROM information_schema.schemata WHERE schema_name = 'auth') THEN
            RAISE EXCEPTION 'schema auth not found';
        END IF;
    END if_auth_schema_exist_test
$$;

DO
$$
    <<if_user_table_exist_test>>
    BEGIN
        IF NOT EXISTS(SELECT 1 FROM pg_tables WHERE schemaname = 'auth' AND tablename = 'user') THEN
            RAISE EXCEPTION 'table auth.user not found';
        END IF;
    END if_user_table_exist_test
$$;

ROLLBACK;
