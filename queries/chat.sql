-- name: UserListByChatID :many
SELECT "user"."id", "user"."username", "user"."given_name", "user"."family_name"
FROM "public"."user"
JOIN "public"."user_chat" ON "user_chat"."user_id" = "user"."id"
WHERE "user_chat"."chat_id"=$1
;

-- name: UserListByChatChannel :many
SELECT "user"."id", "user"."username", "user"."given_name", "user"."family_name"
FROM "public"."user"
JOIN "public"."user_chat" ON "user_chat"."user_id" = "user"."id"
JOIN "public"."chat" ON "chat"."id" = "user_chat"."chat_id"
WHERE "chat"."channel"=$1
;