package usecases

import (
	"github.com/p-jirayusakul/go-clean-arch-template/domain/entities"
)

type AccountsUsecase interface {
	Register(arg entities.AccountsDto) (id string, err error)
	Login(arg entities.AccountsDto) (token string, err error)
}
