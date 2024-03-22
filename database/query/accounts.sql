-- name: CreateAccount :one
INSERT INTO public.accounts(email, password, created_at)
VALUES ($1, $2, NOW()) RETURNING id;

-- name: ListAccounts :many
SELECT id, email, password FROM public.accounts;

-- name: GetAccountByID :one
SELECT id, email, password FROM public.accounts WHERE id = $1 LIMIT 1;

-- name: GetAccountByEmail :one
SELECT id, email, password FROM public.accounts WHERE email = $1 LIMIT 1;

-- name: UpdateAccountPasswordByEmail :exec
UPDATE public.accounts SET password=$2 WHERE email = $1;

-- name: DeleteAccount :exec
DELETE FROM public.accounts WHERE id = $1;

-- name: IsEmailAlreadyExists :one
SELECT CASE
    WHEN count(email) > 0 THEN true
    ELSE false
END AS "isAlreadyExists" FROM public.accounts WHERE email = $1 LIMIT 1;
DELETE FROM public.accounts WHERE id = $1;

-- name: IsAccountAlreadyExists :one
SELECT CASE
    WHEN count(id) > 0 THEN true
    ELSE false
END AS "isAlreadyExists" FROM public.accounts WHERE id = $1 LIMIT 1;