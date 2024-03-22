package http

import (
	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/go-clean-arch-template/domain/usecases"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/repositories/factories"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/usecases/accounts"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/config"
)

type ServerHttpHandler struct {
	Cfg             *config.Config
	AccountsUsecase usecases.AccountsUsecase
}

func NewServerHttpHandler(
	app *echo.Echo,
	config *config.Config,
	dbFactory *factories.DBFactory,

) {
	handler := &ServerHttpHandler{
		Cfg: config,
		AccountsUsecase: accounts.NewAccountsInteractor(
			config,
			dbFactory,
		),
	}

	var baseAPI = "/api/v1"

	authGroup := app.Group(baseAPI + "/auth")
	authGroup.POST("/register", handler.Register)
	authGroup.POST("/login", handler.Login)

}
