-- name: CreateSale :one
INSERT INTO sales (produto_id, quantidade, tempo_valor)
VALUES ($1, $2, $3)
RETURNING id, produto_id, quantidade, tempo_valor, total, data_venda;

