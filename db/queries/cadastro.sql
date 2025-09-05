-- name: CreateCadastro :one
INSERT INTO cadastro (name, cnpj, email, celular, password, status, created_at)
VALUES ($1, $2, $3, $4, $5, $6, now())
RETURNING *;


-- name: GetSellerByCNPJ :one
SELECT id, name, email, celular, password, status, activation_code, created_at
FROM cadastro
WHERE celular= $1;


-- name: UpdateCadastroStatus :exec
UPDATE cadastro
SET status = $2
WHERE activation_code= $1;


-- name: UpdateActivationCode :exec
UPDATE cadastro
SET activation_code = $2,
    status = $3
WHERE celular = $1;
