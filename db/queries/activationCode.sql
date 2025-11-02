-- name: SaveActivationCode :exec
INSERT INTO activation_code (cadastro_id, activation_codes, code, expires_at, status)
VALUES ($1, $2, $3, $4, $5);


-- name: GetActivationCode :one
SELECT a.id AS activation_id,
       a.cadastro_id,
       a.activation_codes,
       a.status AS code_status,
       c.status AS cadastro_status
FROM activation_code a
         JOIN cadastro c ON a.cadastro_id = c.id
WHERE a.activation_codes = $1
  AND a.cadastro_id = $2;
