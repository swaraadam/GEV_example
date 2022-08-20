-- name: Register :one
INSERT INTO m_user (name,password,email,updated_at)
VALUES ($1,$2,$3,$4)
RETURNING *;

-- name: GetUserByUUID :one
SELECT * FROM m_user
WHERE uuid = $1;

-- name: GetUserByEmail :one
SELECT * FROM m_user
WHERE email = $1;

-- name: GetAllUser :many
SELECT * FROM m_user
ORDER BY id;