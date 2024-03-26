package addresses

import (
	"github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/factories"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/worker"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/config"
)

type addressesInteractor struct {
	cfg             *config.Config
	store           factories.Store
	taskDistributor worker.TaskDistributor
}

func NewaddressesInteractor(
	config *config.Config,
	store factories.Store,
	taskDistributor worker.TaskDistributor,
) *addressesInteractor {

	return &addressesInteractor{
		cfg:             config,
		store:           store,
		taskDistributor: taskDistributor,
	}
}
