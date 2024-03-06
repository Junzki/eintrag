CREATE OR REPLACE FUNCTION create_user(
    username character varying,
    password character varying,
    uid uuid default null,
    display_name character varying default null
) RETURNS character varying AS
$uid$
DECLARE
    salt character varying;
BEGIN
    if uid is null then
        uid = gen_random_uuid();
    end if;

    salt = gen_salt('bf');
    password = crypt(password, salt);

    INSERT INTO users (uid, username, password, salt, display_name)
    VALUES (uid, username, password, salt, display_name);

    RETURN uid;
END;

$uid$ LANGUAGE plpgsql;
