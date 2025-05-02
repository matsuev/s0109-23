CREATE TABLE IF NOT EXISTS "public"."user"
(
    "id" uuid NOT NULL,
    "username" character varying(50) NOT NULL DEFAULT '',
    "given_name" character varying(50) NOT NULL DEFAULT '',
    "family_name" character varying(50) NOT NULL DEFAULT '',
    enabled boolean NOT NULL DEFAULT false,
    PRIMARY KEY ("id"),
    UNIQUE ("username")
);

CREATE TABLE IF NOT EXISTS "public"."chat"
(
    "id" bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    "channel" character varying(255) NOT NULL DEFAULT '',
    "title" character varying(255) NOT NULL DEFAULT '',
    PRIMARY KEY ("id"),
    UNIQUE ("channel")
);

CREATE TABLE IF NOT EXISTS "public"."user_chat"
(
    "user_id" uuid NOT NULL,
    "chat_id" bigint NOT NULL,
    "can_publish" boolean NOT NULL DEFAULT false,
    FOREIGN KEY ("user_id")
        REFERENCES "public"."user" ("id") MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID,
    FOREIGN KEY ("chat_id")
        REFERENCES "public"."chat" ("id") MATCH SIMPLE
        ON UPDATE CASCADE
        ON DELETE CASCADE
        NOT VALID
);
