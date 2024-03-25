// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: accounts.sql

package database

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO public.accounts(email, password, created_at)
VALUES ($1, $2, NOW()) RETURNING id
`

type CreateAccountParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg *CreateAccountParams) (string, error) {
	row := q.db.QueryRow(ctx, createAccount, arg.Email, arg.Password)
	var id string
	err := row.Scan(&id)
	return id, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM public.accounts WHERE id = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, id string) error {
	_, err := q.db.Exec(ctx, deleteAccount, id)
	return err
}

const getAccountByEmail = `-- name: GetAccountByEmail :one
SELECT id, email, password FROM public.accounts WHERE email = $1 LIMIT 1
`

type GetAccountByEmailRow struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) GetAccountByEmail(ctx context.Context, email string) (*GetAccountByEmailRow, error) {
	row := q.db.QueryRow(ctx, getAccountByEmail, email)
	var i GetAccountByEmailRow
	err := row.Scan(&i.ID, &i.Email, &i.Password)
	return &i, err
}

const getAccountByID = `-- name: GetAccountByID :one
SELECT id, email, password FROM public.accounts WHERE id = $1 LIMIT 1
`

type GetAccountByIDRow struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) GetAccountByID(ctx context.Context, id string) (*GetAccountByIDRow, error) {
	row := q.db.QueryRow(ctx, getAccountByID, id)
	var i GetAccountByIDRow
	err := row.Scan(&i.ID, &i.Email, &i.Password)
	return &i, err
}

const isAccountAlreadyExists = `-- name: IsAccountAlreadyExists :one
SELECT CASE
    WHEN count(id) > 0 THEN true
    ELSE false
END AS "isAlreadyExists" FROM public.accounts WHERE id = $1 LIMIT 1
`

func (q *Queries) IsAccountAlreadyExists(ctx context.Context, id string) (bool, error) {
	row := q.db.QueryRow(ctx, isAccountAlreadyExists, id)
	var isAlreadyExists bool
	err := row.Scan(&isAlreadyExists)
	return isAlreadyExists, err
}

const isEmailAlreadyExists = `-- name: IsEmailAlreadyExists :one
SELECT CASE
    WHEN count(email) > 0 THEN true
    ELSE false
END AS "isAlreadyExists" FROM public.accounts WHERE email = $1 LIMIT 1
`

func (q *Queries) IsEmailAlreadyExists(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRow(ctx, isEmailAlreadyExists, email)
	var isAlreadyExists bool
	err := row.Scan(&isAlreadyExists)
	return isAlreadyExists, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, email, password FROM public.accounts
`

type ListAccountsRow struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) ListAccounts(ctx context.Context) ([]*ListAccountsRow, error) {
	rows, err := q.db.Query(ctx, listAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ListAccountsRow{}
	for rows.Next() {
		var i ListAccountsRow
		if err := rows.Scan(&i.ID, &i.Email, &i.Password); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccountPasswordByEmail = `-- name: UpdateAccountPasswordByEmail :exec
UPDATE public.accounts SET password=$2 WHERE email = $1
`

type UpdateAccountPasswordByEmailParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (q *Queries) UpdateAccountPasswordByEmail(ctx context.Context, arg *UpdateAccountPasswordByEmailParams) error {
	_, err := q.db.Exec(ctx, updateAccountPasswordByEmail, arg.Email, arg.Password)
	return err
}
