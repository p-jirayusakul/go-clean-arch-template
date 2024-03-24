package http

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/go-clean-arch-template/domain/usecases"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/factories"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/usecases/accounts"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/usecases/addresses"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/common"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/config"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/middleware"
)

type ServerHttpHandler struct {
	Cfg              *config.Config
	AccountsUsecase  usecases.AccountsUsecase
	AddressesUsecase usecases.AddressesUsecase
}

func NewServerHttpHandler(
	app *echo.Echo,
	cfg *config.Config,
	dbFactory *factories.DBFactory,

) {
	handler := &ServerHttpHandler{
		Cfg: cfg,
		AccountsUsecase: accounts.NewAccountsInteractor(
			cfg,
			dbFactory,
		),
		AddressesUsecase: addresses.NewaddressesInteractor(
			cfg,
			dbFactory,
		),
	}

	var baseAPI = "/api/v1"

	// auth
	authGroup := app.Group(baseAPI + "/auth")
	authGroup.POST("/register", handler.Register)
	authGroup.POST("/login", handler.Login)

	// profile
	addressesGroup := app.Group(baseAPI + "/profile")
	addressesGroup.Use(middleware.ConfigJWT(cfg.JWT_SECRET))
	addressesGroup.POST("/addresses", handler.CreateAddresses)
	addressesGroup.GET("/addresses", handler.ListAddresses)
	addressesGroup.PUT("/addresses/:id", handler.UpdateAddresses)
	addressesGroup.DELETE("/addresses/:id", handler.DeleteAddresses)
}

// utils function
func (s *ServerHttpHandler) GetTokenID(c echo.Context) error {
	isAlreadyExists, err := s.AccountsUsecase.IsAccountAlreadyExists(c.Get("accountsID").(string))
	if err != nil {
		return err
	}

	if !isAlreadyExists {
		return fmt.Errorf(common.ErrAccountIsInvalid.Error())
	}

	return nil
}
