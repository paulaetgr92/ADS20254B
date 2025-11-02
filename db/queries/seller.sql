-- name: CreateSeller :one
INSERT INTO cadastro (
    activation_code,
    name,
    email,
    password,
    cpf,
    cnpj,
    celular,
    status,
    created_at
) VALUES (
             $1, $2, $3, $4, $5, $6, $7, $8,NOW()
         )
RETURNING id, name, email, status,  activation_code, celular;
