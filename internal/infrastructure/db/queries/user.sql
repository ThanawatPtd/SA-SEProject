-- name: GetAllUsers :many
SELECT
    id,
    fname AS name,
    email,
    created_at,
    updated_at
FROM "user";

-- name: GetUserByID :one
SELECT
    id,
    fname AS name,
    email,
    created_at,
    updated_at
FROM "user"
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO "user" (
    id, email, fname, lname, password, phone_number, address, created_at, updated_at
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, NOW(), NOW()
)
RETURNING id, fname AS name, email, created_at, updated_at;

-- name: UpdateUser :one
UPDATE "user"
SET
    email = $2,
    fname = $3,
    lname = $4,
    password = $5,
    phone_number = $6,
    address = $7,
    updated_at = NOW()
WHERE id = $1
RETURNING id, fname AS name, email, created_at, updated_at;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;
