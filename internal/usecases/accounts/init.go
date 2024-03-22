package accounts

import (
	"github.com/p-jirayusakul/go-clean-arch-template/domain/repositories"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/factories"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/config"
)

type accountsInteractor struct {
	cfg          *config.Config
	accountsRepo repositories.AccountsRepository
}

func NewAccountsInteractor(
	config *config.Config,
	dbFactory *factories.DBFactory,
) *accountsInteractor {

	return &accountsInteractor{
		cfg:          config,
		accountsRepo: dbFactory.AccountsRepo,
	}
}
