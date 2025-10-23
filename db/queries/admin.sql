
-- name: CreateLoginAdmin :one
SELECT name, email, password
FROM admin
WHERE email = $1
  AND password = $2
LIMIT 1;


-- name: CreateAdmin :one
INSERT INTO admin (name, email, cnpj,password)
VALUES (
           $1,
           $2,
        $3,
    $4
       )
RETURNING *;