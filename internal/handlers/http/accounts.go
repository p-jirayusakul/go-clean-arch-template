package http

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/go-clean-arch-template/domain/entities"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/handlers/http/request"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/handlers/http/response"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/common"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/utils"
)

// Register
// @Summary      Register By email and password
// @Description  register
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param request body request.RegisterRequest true "body request"
// @Success      201  {object}  utils.SuccessResponse
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /api/v1/core/auth/register [post]
func (s *ServerHttpHandler) Register(c echo.Context) (err error) {

	// pare json
	body := new(request.RegisterRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// Logic
	arg := entities.AccountsDto{
		Email:    body.Email,
		Password: body.Password,
	}

	_, err = s.AccountsUsecase.Register(arg)
	if err != nil {
		if errors.Is(err, common.ErrEmailIsAlreadyExists) {
			return utils.RespondWithError(http.StatusBadRequest, common.ErrEmailIsAlreadyExists.Error())
		}
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	// Response
	var payload interface{}
	message := "registration completed"
	return utils.RespondWithJSON(c, http.StatusCreated, message, payload)
}

// Login
// @Summary      Login By email and password
// @Description  register
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param request body request.LoginRequest true "body request"
// @Success      200  {object}  utils.SuccessResponse.Data{data=response.LoginResponse}
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /api/v1/core/auth/login [post]
func (s *ServerHttpHandler) Login(c echo.Context) (err error) {

	// pare json
	body := new(request.LoginRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// Logic
	arg := entities.AccountsDto{
		Email:    body.Email,
		Password: body.Password,
	}

	result, err := s.AccountsUsecase.Login(arg)
	if err != nil {
		if errors.Is(err, common.ErrLoginFail) {
			return utils.RespondWithError(http.StatusUnauthorized, common.ErrLoginFail.Error())
		}
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	// Response
	payload := response.LoginResponse{
		Token: result,
	}
	message := "login completed"
	return utils.RespondWithJSON(c, http.StatusOK, message, payload)
}
