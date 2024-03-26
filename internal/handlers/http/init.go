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

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/p-jirayusakul/go-clean-arch-template/docs"
)

type ServerHttpHandler struct {
	Cfg              *config.Config
	AccountsUsecase  usecases.AccountsUsecase
	AddressesUsecase usecases.AddressesUsecase
}

func NewServerHttpHandler(
	app *echo.Echo,
	cfg *config.Config,
	store factories.Store,

) *ServerHttpHandler {
	handler := &ServerHttpHandler{
		Cfg: cfg,
		AccountsUsecase: accounts.NewAccountsInteractor(
			cfg,
			store,
		),
		AddressesUsecase: addresses.NewaddressesInteractor(
			cfg,
			store,
		),
	}

	app.GET(common.DOCS_URL+"/*", echoSwagger.WrapHandler)

	// auth
	authGroup := app.Group(common.BASE_URL + "/auth")
	authGroup.POST("/register", handler.Register)
	authGroup.POST("/login", handler.Login)

	// profile
	addressesGroup := app.Group(common.BASE_URL + "/profile")
	addressesGroup.Use(middleware.ConfigJWT(cfg.JWT_SECRET))
	addressesGroup.POST("/addresses", handler.CreateAddresses)
	addressesGroup.GET("/addresses/me", handler.ListAddresses)
	addressesGroup.GET("/addresses", handler.SearchAddresses)
	addressesGroup.PUT("/addresses/:id", handler.UpdateAddresses)
	addressesGroup.DELETE("/addresses/:id", handler.DeleteAddresses)

	return handler
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
