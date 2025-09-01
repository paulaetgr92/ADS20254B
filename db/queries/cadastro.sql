-- name: CreateCadastro :one
INSERT INTO cadastro (name, email, password, AcessID,created_at)
    VALUES ($1, $2, $3,$4,now())
RETURNING *;
