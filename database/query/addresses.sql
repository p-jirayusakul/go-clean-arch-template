-- name: CreateAddresses :one
INSERT INTO public.addresses(street_address, city, state_province, postal_code, country, accounts_id)
VALUES ($1, $2, $3, $4, $5, sqlc.arg(accounts_id)::text) RETURNING id;

-- name: ListAddresses :many
SELECT id, street_address as "streetAddress", city, state_province as "stateProvince", postal_code as "postalCode", country FROM public.addresses;

-- name: ListAddressesByAccountId :many
SELECT id, street_address as "streetAddress", city, state_province as "stateProvince", postal_code as "postalCode", country FROM public.addresses WHERE accounts_id = sqlc.arg(accounts_id)::text;

-- name: GetAddressById :one
SELECT id, street_address as "streetAddress", city, state_province as "stateProvince", postal_code as "postalCode", country FROM public.addresses WHERE id = $1 LIMIT 1;

-- name: UpdateAddressById :exec
UPDATE public.addresses
	SET updated_at = NOW(), street_address=$2, city=$3, state_province=$4, postal_code=$5, country=$6 
WHERE id = $1 AND accounts_id = sqlc.arg(accounts_id)::text;

-- name: DeleteAddressesById :exec
DELETE FROM public.addresses WHERE id = $1;

-- name: IsAddressesAlreadyExists :one
SELECT CASE
    WHEN count(id) > 0 THEN true
    ELSE false
END AS "isAlreadyExists" FROM public.addresses WHERE id = $1 AND accounts_id = sqlc.arg(accounts_id)::text LIMIT 1;