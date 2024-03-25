package accounts

import (
	"github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/factories"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/config"
)

type accountsInteractor struct {
	cfg       *config.Config
	dbFactory factories.DBFactory
}

func NewAccountsInteractor(
	config *config.Config,
	dbFactory factories.DBFactory,
) *accountsInteractor {

	return &accountsInteractor{
		cfg:       config,
		dbFactory: dbFactory,
	}
}
