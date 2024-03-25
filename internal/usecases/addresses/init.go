package addresses

import (
	"github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/factories"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/config"
)

type addressesInteractor struct {
	cfg   *config.Config
	store factories.Store
}

func NewaddressesInteractor(
	config *config.Config,
	store factories.Store,
) *addressesInteractor {

	return &addressesInteractor{
		cfg:   config,
		store: store,
	}
}
