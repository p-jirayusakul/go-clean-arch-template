package addresses

import (
	"github.com/p-jirayusakul/go-clean-arch-template/domain/repositories"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/factories"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/config"
)

type addressesInteractor struct {
	cfg           *config.Config
	addressesRepo repositories.AddressesRepository
}

func NewaddressesInteractor(
	config *config.Config,
	dbFactory *factories.DBFactory,
) *addressesInteractor {

	return &addressesInteractor{
		cfg:           config,
		addressesRepo: dbFactory.AddressesRepo,
	}
}
