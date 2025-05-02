
-- name: UserCanSubscribe :one
SELECT "id" FROM "public"."user_chat"
JOIN "public"."chat" ON "chat"."id" = "user_chat"."chat_id"
WHERE "user_chat"."user_id"=$1 AND "chat"."channel"=$2
;

-- name: UserCanPublish :one
SELECT "can_publish" FROM "public"."user_chat"
JOIN "public"."chat" ON "chat"."id" = "user_chat"."chat_id"
WHERE "user_chat"."user_id"=$1 AND "chat"."channel"=$2
;

