-- name: CreateCadastro :one
INSERT INTO cadastro (name, cpf, cnpj, email, celular, password, status, activation_code)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetSellerByCNPJ :one
SELECT id, name, email, celular, password, status, activation_code, created_at
FROM cadastro
WHERE celular= $1;


-- name: GetCadastroByActivationCode :one
SELECT id, name, email, celular, status, activation_code
FROM cadastro
WHERE activation_code = $1 and id=$2;


-- name: UpdateCadastroStatus :exec
UPDATE cadastro
SET status = $2,
    updated_at = NOW()
WHERE id = $1;
