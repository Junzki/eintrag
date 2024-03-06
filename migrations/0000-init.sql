
CREATE TABLE "users" (
    id serial primary key,
    uid uuid not null,
    username character varying not null,
    display_name character varying default null,
    password character varying default null,
    salt character varying default null,
    created_at timestamp with time zone not null default current_timestamp,
    updated_at timestamp with time zone not null default current_timestamp,
    last_seen timestamp with time zone not null default current_timestamp,

    unique (uid, username)
);
