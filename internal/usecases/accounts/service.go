package accounts

import (
	"context"
	"errors"
	"time"

	"github.com/p-jirayusakul/go-clean-arch-template/domain/entities"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/common"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/middleware"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/utils"
)

func (x *accountsInteractor) Register(arg entities.AccountsDto) (id string, err error) {
	ctx := context.Background()

	isEmailAlready, err := x.accountsRepo.IsEmailAlreadyExists(ctx, arg.Email)
	if err != nil {
		return
	}

	if isEmailAlready {
		return "", common.ErrEmailIsAlreadyExists
	}

	hashedPassword, err := utils.HashPassword(arg.Password)
	if err != nil {
		return
	}

	params := entities.AccountsDto{
		Email:    arg.Email,
		Password: hashedPassword,
	}

	id, err = x.accountsRepo.CraeteAccount(ctx, params)
	if err != nil {
		return
	}

	return
}

func (x *accountsInteractor) Login(arg entities.AccountsDto) (token string, err error) {
	ctx := context.Background()

	account, err := x.accountsRepo.GetAccountByEmail(ctx, arg.Email)
	if err != nil {
		if errors.Is(err, common.ErrDBNoRows) {
			return "", common.ErrLoginFail
		}
		return
	}

	err = utils.CheckPassword(arg.Password, account.Password)
	if err != nil {
		return "", common.ErrLoginFail
	}

	token, err = middleware.CreateToken(middleware.CreateTokenDTO{
		UserID:    account.ID,
		Secret:    x.cfg.JWT_SECRET,
		ExpiresAt: time.Now().Add(time.Hour * 72),
	})

	if err != nil {
		return
	}

	return
}

func (x *accountsInteractor) IsAccountAlreadyExists(arg string) (isAlreadyExists bool, err error) {
	ctx := context.Background()

	isAlreadyExists, err = x.accountsRepo.IsAccountAlreadyExists(ctx, arg)
	if err != nil {
		if errors.Is(err, common.ErrDBNoRows) {
			return false, common.ErrLoginFail
		}
		return
	}

	return
}