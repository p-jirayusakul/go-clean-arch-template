package accounts

import (
	"github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/factories"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/config"
)

type accountsInteractor struct {
	cfg   *config.Config
	store factories.Store
}

func NewAccountsInteractor(
	config *config.Config,
	store factories.Store,
) *accountsInteractor {

	return &accountsInteractor{
		cfg:   config,
		store: store,
	}
}
