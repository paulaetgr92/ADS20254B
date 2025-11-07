-- name: SaveActivationCode :one
INSERT INTO activation_code (
    cadastro_id,
    activation_code,
    code,
    expires_at,
    status
)
VALUES ($1, $2, $3, $4, $5)
RETURNING cadastro_id, activation_code, code;


-- name: GetActivationCode :one
SELECT a.id AS activation_code,
       a.cadastro_id,
       a.activation_code,
       a.status AS code_status,
       a.expires_at,
       c.status AS cadastro_status
FROM activation_code a
         JOIN cadastro c ON a.cadastro_id = c.id
WHERE a.cadastro_id = $1
ORDER BY a.created_at DESC
LIMIT 1;


-- name: VerifyActivationCode :one
SELECT
    id,
    cadastro_id,
    code

FROM activation_code
WHERE cadastro_id = $1
  AND code = $2
LIMIT 1;
