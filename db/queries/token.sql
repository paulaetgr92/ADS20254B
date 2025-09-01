-- name: GetUserTokensHist :one
SELECT id, "token", expires_at, active
FROM public.token
WHERE user_id=$1 AND
    access_id=$2 ;

-- name: CreateTokenHist :exec
INSERT INTO public.token
(id, user_id, "token", expires_at, active, created_at,access_id)
VALUES(nextval('token'::regclass), $1, $2, $3,true, now(), $4);