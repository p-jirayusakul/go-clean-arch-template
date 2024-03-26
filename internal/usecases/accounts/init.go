package accounts

import (
	"github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/factories"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/worker"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/config"
)

type accountsInteractor struct {
	cfg             *config.Config
	store           factories.Store
	taskDistributor worker.TaskDistributor
}

func NewAccountsInteractor(
	config *config.Config,
	store factories.Store,
	taskDistributor worker.TaskDistributor,
) *accountsInteractor {

	return &accountsInteractor{
		cfg:             config,
		store:           store,
		taskDistributor: taskDistributor,
	}
}
