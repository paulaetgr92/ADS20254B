-- name: CreateSeller :one
INSERT INTO cadastro (
    name,
    email,
                      activation_code,
    password,
    cpf,
    cnpj,
    celular,
    status,
    created_at
) VALUES (
             $1, $2, $3, $4, $5, $6, $7,$8, NOW()
         )
RETURNING id AS cadastro_id,activation_code, status;
