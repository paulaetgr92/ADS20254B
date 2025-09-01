
-- name: CreateLogin :one
INSERT INTO login (email, password)
VALUES ($1, $2)
RETURNING *;


-- name: GetLogin :one
SELECT email, password
FROM login
WHERE email = $1;
