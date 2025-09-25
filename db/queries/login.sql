-- name: GetLogin :one
SELECT email, password, activation_code,status
FROM cadastro
WHERE email = $1;


-- name: CreateLogin :one
INSERT INTO cadastro (email, password, activation_code, status)
VALUES ($1, $2, $3, $4)
RETURNING id, email, password, activation_code, status, created_at;




