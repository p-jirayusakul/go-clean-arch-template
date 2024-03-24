package repositories

import (
	"context"

	"github.com/p-jirayusakul/go-clean-arch-template/domain/entities"
)

type AccountsRepository interface {
	CraeteAccount(ctx context.Context, p entities.AccountsDto) (string, error)
	IsEmailAlreadyExists(ctx context.Context, p string) (bool, error)
	IsAccountAlreadyExists(ctx context.Context, p string) (bool, error)
	GetAccountByEmail(ctx context.Context, p string) (entities.Accounts, error)
}
