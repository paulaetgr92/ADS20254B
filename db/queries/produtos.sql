-- name: GetProdutoById :one
SELECT categoria, tamanho, cores, localizacao
FROM produto
WHERE id_roupa = $1;

-- name: CreateProduct :one
INSERT INTO produto (categoria, tamanho, cores, tempo_valor, status, localizacao)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id_roupa;

-- name: GetProdutoByDisponibilidade :many
SELECT categoria, tamanho, cores
FROM produto
WHERE id_roupa = $1 AND status = $2 AND localizacao = $3;

-- name: AtualizarProdutoByID :one
UPDATE produto
SET categoria = $1,
    tamanho = $2,
    cores = $3,
    tempo_valor = $4,
    status = $5,
    localizacao = $6
WHERE id_roupa = $7
RETURNING *;


-- name: DeleteProdutoByID :one
UPDATE produto
SET categoria = $1,
    tamanho = $2,
    cores = $3,
    tempo_valor = $4,
    status = $5,
    localizacao = $6
WHERE id_roupa = $7 and status = true
RETURNING *;


