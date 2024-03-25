package http

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/go-clean-arch-template/domain/entities"
	"github.com/p-jirayusakul/go-clean-arch-template/internal/handlers/http/request"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/common"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/utils"
)

func (s *ServerHttpHandler) CreateAddresses(c echo.Context) (err error) {

	// pare json
	body := new(request.CreateAddressesRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	err = s.GetTokenID(c)
	if err != nil {
		return utils.RespondWithError(http.StatusUnauthorized, err.Error())
	}

	accontID := c.Get("accountsID").(string)

	// Logic
	arg := entities.AddressesDto{
		StreetAddress: body.Address,
		City:          body.City,
		StateProvince: body.Province,
		PostalCode:    body.PostalCode,
		Country:       body.Country,
		AccountsID:    accontID,
	}

	_, err = s.AddressesUsecase.CreateAddresses(arg)
	if err != nil {
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	// Response
	var payload interface{}
	message := "create addresses completed"
	return utils.RespondWithJSON(c, http.StatusCreated, message, payload)
}

func (s *ServerHttpHandler) ListAddresses(c echo.Context) (err error) {

	err = s.GetTokenID(c)
	if err != nil {
		return utils.RespondWithError(http.StatusUnauthorized, err.Error())
	}

	accontID := c.Get("accountsID").(string)

	result, err := s.AddressesUsecase.ListAddressesAddresses(accontID)
	if err != nil {
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	// Response
	message := "get addresses completed"
	return utils.RespondWithJSON(c, http.StatusOK, message, result)
}

func (s *ServerHttpHandler) UpdateAddresses(c echo.Context) (err error) {

	// pare json
	body := new(request.UpdateAddressesRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	err = s.GetTokenID(c)
	if err != nil {
		return utils.RespondWithError(http.StatusUnauthorized, err.Error())
	}

	accontID := c.Get("accountsID").(string)

	// Logic
	arg := entities.AddressesDto{
		ID:            body.ID,
		StreetAddress: body.Address,
		City:          body.City,
		StateProvince: body.Province,
		PostalCode:    body.PostalCode,
		Country:       body.Country,
		AccountsID:    accontID,
	}

	err = s.AddressesUsecase.UpdateAddresses(arg)
	if err != nil {
		if errors.Is(err, common.ErrDataNotFound) {
			return utils.RespondWithError(http.StatusNotFound, err.Error())
		}
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	// Response
	var payload interface{}
	message := "update addresses completed"
	return utils.RespondWithJSON(c, http.StatusOK, message, payload)
}

func (s *ServerHttpHandler) DeleteAddresses(c echo.Context) (err error) {

	// pare json
	body := new(request.DeleteAddressesRequest)
	if err := c.Bind(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	// validate DTO
	if err = c.Validate(body); err != nil {
		return utils.RespondWithError(http.StatusBadRequest, err.Error())
	}

	err = s.GetTokenID(c)
	if err != nil {
		return utils.RespondWithError(http.StatusUnauthorized, err.Error())
	}

	accontID := c.Get("accountsID").(string)

	// Logic
	arg := entities.AddressesDto{
		ID:         body.ID,
		AccountsID: accontID,
	}

	err = s.AddressesUsecase.DeleteAddresses(arg)
	if err != nil {
		if errors.Is(err, common.ErrDataNotFound) {
			return utils.RespondWithError(http.StatusNotFound, err.Error())
		}
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	// Response
	var payload interface{}
	message := "deleted addresses completed"
	return utils.RespondWithJSON(c, http.StatusNoContent, message, payload)
}
