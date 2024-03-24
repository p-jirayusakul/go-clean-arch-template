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

func (x *AccountsRepository) CraeteAccount(ctx context.Context, account entities.AccountsDto) (result string, err error) {
	params := database.CreateAccountParams{
		Email:    account.Email,
		Password: account.Password,
	}

	result, err = x.db.CreateAccount(ctx, params)
	if err != nil {
		return
	}

	return
}

func (x *AccountsRepository) IsEmailAlreadyExists(ctx context.Context, email string) (result bool, err error) {
	result, err = x.db.IsEmailAlreadyExists(ctx, email)
	if err != nil {
		return
	}

	return
}

func (x *AccountsRepository) IsAccountAlreadyExists(ctx context.Context, accountsID string) (result bool, err error) {
	result, err = x.db.IsAccountAlreadyExists(ctx, accountsID)
	if err != nil {
		return
	}

	return
}

func (x *AccountsRepository) GetAccountByEmail(ctx context.Context, email string) (result entities.Accounts, err error) {
	r, err := x.db.GetAccountByEmail(ctx, email)
	if err != nil {
		return
	}

	result = entities.Accounts{
		ID:       r.ID,
		Email:    r.Email,
		Password: r.Password,
	}

	return
}
