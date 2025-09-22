-- name: CreateLogin :one
INSERT INTO cadastro (email, password,activation_code)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetLogin :one
SELECT email, password, activation_code
FROM cadastro
WHERE email = $1;

-- name: UpdateAuthenticationCode :exec
UPDATE cadastro
SET activation_code = $2
WHERE email = $1;
