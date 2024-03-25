package addresses

import (
	"github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/factories"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/config"
)

type addressesInteractor struct {
	cfg       *config.Config
	dbFactory factories.DBFactory
}

func NewaddressesInteractor(
	config *config.Config,
	dbFactory factories.DBFactory,
) *addressesInteractor {

	return &addressesInteractor{
		cfg:       config,
		dbFactory: dbFactory,
	}
}
