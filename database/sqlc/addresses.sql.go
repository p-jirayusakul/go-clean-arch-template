// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: addresses.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAddresses = `-- name: CreateAddresses :one
INSERT INTO public.addresses(street_address, city, state_province, postal_code, country, accounts_id)
VALUES ($1, $2, $3, $4, $5, $6::text) RETURNING id
`

type CreateAddressesParams struct {
	StreetAddress pgtype.Text `json:"street_address"`
	City          string      `json:"city"`
	StateProvince string      `json:"state_province"`
	PostalCode    string      `json:"postal_code"`
	Country       string      `json:"country"`
	AccountsID    string      `json:"accounts_id"`
}

func (q *Queries) CreateAddresses(ctx context.Context, arg CreateAddressesParams) (string, error) {
	row := q.db.QueryRow(ctx, createAddresses,
		arg.StreetAddress,
		arg.City,
		arg.StateProvince,
		arg.PostalCode,
		arg.Country,
		arg.AccountsID,
	)
	var id string
	err := row.Scan(&id)
	return id, err
}

const deleteAddressesById = `-- name: DeleteAddressesById :exec
DELETE FROM public.addresses WHERE id = $1
`

func (q *Queries) DeleteAddressesById(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteAddressesById, id)
	return err
}

const getAddressById = `-- name: GetAddressById :one
SELECT id, street_address, city, state_province, postal_code, country, accounts_id FROM public.addresses WHERE id = $1 LIMIT 1
`

type GetAddressByIdRow struct {
	ID            string      `json:"id"`
	StreetAddress pgtype.Text `json:"street_address"`
	City          string      `json:"city"`
	StateProvince string      `json:"state_province"`
	PostalCode    string      `json:"postal_code"`
	Country       string      `json:"country"`
	AccountsID    pgtype.Text `json:"accounts_id"`
}

func (q *Queries) GetAddressById(ctx context.Context, id string) (GetAddressByIdRow, error) {
	row := q.db.QueryRow(ctx, getAddressById, id)
	var i GetAddressByIdRow
	err := row.Scan(
		&i.ID,
		&i.StreetAddress,
		&i.City,
		&i.StateProvince,
		&i.PostalCode,
		&i.Country,
		&i.AccountsID,
	)
	return i, err
}

const isAddressesAlreadyExists = `-- name: IsAddressesAlreadyExists :one
SELECT CASE
    WHEN count(id) > 0 THEN true
    ELSE false
END AS "isAlreadyExists" FROM public.addresses WHERE id = $1 AND accounts_id = $2::text LIMIT 1
`

type IsAddressesAlreadyExistsParams struct {
	ID         string `json:"id"`
	AccountsID string `json:"accounts_id"`
}

func (q *Queries) IsAddressesAlreadyExists(ctx context.Context, arg IsAddressesAlreadyExistsParams) (bool, error) {
	row := q.db.QueryRow(ctx, isAddressesAlreadyExists, arg.ID, arg.AccountsID)
	var isAlreadyExists bool
	err := row.Scan(&isAlreadyExists)
	return isAlreadyExists, err
}

const listAddresses = `-- name: ListAddresses :many
SELECT id, street_address, city, state_province, postal_code, country, accounts_id FROM public.addresses
`

type ListAddressesRow struct {
	ID            string      `json:"id"`
	StreetAddress pgtype.Text `json:"street_address"`
	City          string      `json:"city"`
	StateProvince string      `json:"state_province"`
	PostalCode    string      `json:"postal_code"`
	Country       string      `json:"country"`
	AccountsID    pgtype.Text `json:"accounts_id"`
}

func (q *Queries) ListAddresses(ctx context.Context) ([]ListAddressesRow, error) {
	rows, err := q.db.Query(ctx, listAddresses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListAddressesRow{}
	for rows.Next() {
		var i ListAddressesRow
		if err := rows.Scan(
			&i.ID,
			&i.StreetAddress,
			&i.City,
			&i.StateProvince,
			&i.PostalCode,
			&i.Country,
			&i.AccountsID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAddressesByAccountId = `-- name: ListAddressesByAccountId :many
SELECT id, street_address, city, state_province, postal_code, country, accounts_id FROM public.addresses WHERE accounts_id = $1::text
`

type ListAddressesByAccountIdRow struct {
	ID            string      `json:"id"`
	StreetAddress pgtype.Text `json:"street_address"`
	City          string      `json:"city"`
	StateProvince string      `json:"state_province"`
	PostalCode    string      `json:"postal_code"`
	Country       string      `json:"country"`
	AccountsID    pgtype.Text `json:"accounts_id"`
}

func (q *Queries) ListAddressesByAccountId(ctx context.Context, accountsID string) ([]ListAddressesByAccountIdRow, error) {
	rows, err := q.db.Query(ctx, listAddressesByAccountId, accountsID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListAddressesByAccountIdRow{}
	for rows.Next() {
		var i ListAddressesByAccountIdRow
		if err := rows.Scan(
			&i.ID,
			&i.StreetAddress,
			&i.City,
			&i.StateProvince,
			&i.PostalCode,
			&i.Country,
			&i.AccountsID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAddressById = `-- name: UpdateAddressById :exec
UPDATE public.addresses
	SET updated_at = NOW(), street_address=$2, city=$3, state_province=$4, postal_code=$5, country=$6 
WHERE id = $1 AND accounts_id = $7::text
`

type UpdateAddressByIdParams struct {
	ID            string      `json:"id"`
	StreetAddress pgtype.Text `json:"street_address"`
	City          string      `json:"city"`
	StateProvince string      `json:"state_province"`
	PostalCode    string      `json:"postal_code"`
	Country       string      `json:"country"`
	AccountsID    string      `json:"accounts_id"`
}

func (q *Queries) UpdateAddressById(ctx context.Context, arg UpdateAddressByIdParams) error {
	_, err := q.db.Exec(ctx, updateAddressById,
		arg.ID,
		arg.StreetAddress,
		arg.City,
		arg.StateProvince,
		arg.PostalCode,
		arg.Country,
		arg.AccountsID,
	)
	return err
}