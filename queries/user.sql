
-- name: UserByID :one
SELECT "id", "username", "given_name", "family_name"
FROM "public"."user"
WHERE "id"=$1
;

-- name: UserCreate :exec
INSERT INTO "public"."user"
("id", "username", "given_name", "family_name")
VALUES ($1, $2, $3, $4)
;