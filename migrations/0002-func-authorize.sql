CREATE OR REPLACE FUNCTION secure_compare(
    a character varying,
    b character varying
)
    RETURNS boolean AS
$matched$
DECLARE
    matched boolean;
    size    integer;
    chr_a   character default null;
    chr_b   character default null;
BEGIN
    size = (CASE WHEN LENGTH(a) > LENGTH(b) THEN LENGTH(a) ELSE LENGTH(b) END);

    matched = TRUE;
    FOR index IN 1..size
        LOOP
            chr_a = substr(a, index, 1);
            chr_b = substr(b, index, 1);

            IF chr_a IS NULL OR chr_b IS NULL THEN
                matched = FALSE;
                continue;
            END IF;

            IF chr_a <> chr_b THEN
                matched = FALSE;
                CONTINUE;
            END IF;

        END LOOP;

    RETURN matched;
END
$matched$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION authorize_user_with_password(
    input_username character varying,
    input_password character varying
)
    RETURNS TABLE
            (
                uid            uuid,
                found_username character varying,
                display_name   character varying
            )
AS
$user_info$
DECLARE
    stored_password character varying;
    salt            character varying;
BEGIN
    SELECT u.uid,
           u.username,
           u.password as "stored_password",
           u.salt,
           u.display_name
    FROM "users" u
    WHERE u.username = input_username
    LIMIT 1
    INTO uid, found_username, stored_password, salt, display_name;

    if uid IS NULL OR stored_password IS NULL THEN
        raise notice 'user not found';
        RETURN QUERY
            SELECT NULL::uuid AS uid,
                   NULL AS found_username,
                   NULL AS display_name;
        RETURN;
    END IF;

    raise notice 'Found: uid: %, %', uid, uid is null;

    IF secure_compare(crypt(input_password, salt), stored_password) THEN
        RETURN QUERY
            SELECT uid AS uid,
                   found_username AS found_username,
                   display_name AS display_name;
    ELSE
        RETURN QUERY
            SELECT NULL::uuid AS uid,
                   NULL AS found_username,
                   NULL AS display_name;
    END IF;
    RETURN;
END
$user_info$ LANGUAGE plpgsql;
