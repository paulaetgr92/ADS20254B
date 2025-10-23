-- name: GetProdutoById :one
SELECT id_roupa, categoria, tamanho, cores, localizacao, tempo_valor, status, imagem_url
FROM produto
WHERE id_roupa = $1;

-- name: CreateProduct :one
INSERT INTO produto (categoria, tamanho, cores, tempo_valor, status, localizacao, imagem_url)
VALUES ($1, $2, $3, $4, $5, $6, $7)
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
    localizacao = $6,
    imagem_url = $7
WHERE id_roupa = $8
RETURNING *;


-- name: InativarProdutoByID :exec
UPDATE produto
SET status = 'inativo'
WHERE id_roupa = $1;



-- name: ListProducts :many
SELECT *
FROM produto;




-- name: GetAllProdutos :many
SELECT * FROM produto ORDER BY id_roupa DESC;