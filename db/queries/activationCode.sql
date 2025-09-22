-- name: SaveActivationCode :exec
INSERT INTO activation_codes (cadastro_id, code, expires_at)
VALUES ($1, $2, $3);
