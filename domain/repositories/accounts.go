package repositories

import (
	"context"

	"github.com/p-jirayusakul/go-clean-arch-template/domain/entities"
)

type AccountsRepository interface {
	CraeteAccount(ctx context.Context, account entities.AccountsDto) (result string, err error)
	IsEmailAlreadyExists(ctx context.Context, email string) (result bool, err error)
	IsAccountAlreadyExists(ctx context.Context, accountsID string) (result bool, err error)
	GetAccountByEmail(ctx context.Context, email string) (result entities.Accounts, err error)
}
