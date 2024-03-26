package accounts

import (
	"context"
	"errors"
	"time"

	"github.com/hibiken/asynq"
	database "github.com/p-jirayusakul/go-clean-arch-template/database/sqlc"
	"github.com/p-jirayusakul/go-clean-arch-template/domain/entities"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/worker"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/common"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/middleware"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/utils"
)

func (s *accountsInteractor) Register(arg entities.AccountsDto) (id string, err error) {
	ctx := context.Background()

	isEmailAlready, err := s.store.IsEmailAlreadyExists(ctx, arg.Email)
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

	params := database.CreateAccountParams{
		Email:    arg.Email,
		Password: hashedPassword,
	}

	id, err = s.store.CreateAccount(ctx, &params)
	if err != nil {
		return
	}

	taskPayload := &worker.PayloadSendVerifyEmail{
		Email: arg.Email,
	}

	opts := []asynq.Option{
		asynq.MaxRetry(10),
		asynq.ProcessIn(5 * time.Second),
		asynq.Queue(worker.QueueCritical),
	}

	s.taskDistributor.DistributeTaskSendVerifyEmail(ctx, taskPayload, opts...)

	return
}

func (s *accountsInteractor) Login(arg entities.AccountsDto) (token string, err error) {
	ctx := context.Background()

	account, err := s.store.GetAccountByEmail(ctx, arg.Email)
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
		Secret:    s.cfg.JWT_SECRET,
		ExpiresAt: time.Now().Add(time.Hour * 72),
	})

	if err != nil {
		return
	}

	return
}

func (s *accountsInteractor) IsAccountAlreadyExists(arg string) (isAlreadyExists bool, err error) {
	ctx := context.Background()

	isAlreadyExists, err = s.store.IsAccountAlreadyExists(ctx, arg)
	if err != nil {
		return
	}

	return
}
