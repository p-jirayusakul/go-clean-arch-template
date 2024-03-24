package db

import (
	"context"

	database "github.com/p-jirayusakul/go-clean-arch-template/database/sqlc"
	"github.com/p-jirayusakul/go-clean-arch-template/domain/entities"
)

type AccountsRepository struct {
	db *database.Queries
}

func NewAccountsRepository(db *database.Queries) AccountsRepository {
	return AccountsRepository{db: db}
}

func (x *AccountsRepository) CraeteAccount(ctx context.Context, p entities.AccountsDto) (string, error) {
	params := database.CreateAccountParams{
		Email:    p.Email,
		Password: p.Password,
	}

	r, err := x.db.CreateAccount(ctx, params)
	if err != nil {
		return "", err
	}

	return r, nil
}

func (x *AccountsRepository) IsEmailAlreadyExists(ctx context.Context, p string) (bool, error) {
	r, err := x.db.IsEmailAlreadyExists(ctx, p)
	if err != nil {
		return false, err
	}

	return r, nil
}

func (x *AccountsRepository) IsAccountAlreadyExists(ctx context.Context, p string) (bool, error) {
	r, err := x.db.IsAccountAlreadyExists(ctx, p)
	if err != nil {
		return false, err
	}

	return r, nil
}

func (x *AccountsRepository) GetAccountByEmail(ctx context.Context, p string) (entities.Accounts, error) {
	r, err := x.db.GetAccountByEmail(ctx, p)
	if err != nil {
		return entities.Accounts{}, err
	}

	e := entities.Accounts{
		ID:       r.ID,
		Email:    r.Email,
		Password: r.Password,
	}

	return e, nil
}
