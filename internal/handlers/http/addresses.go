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

// Address
// @Summary      Create Address
// @Description  register
// @Tags         profile
// @Accept       json
// @Produce      json
// @Param request body request.CreateAddressesRequest true "body request"
// @Success      201  {object}  utils.SuccessResponse
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /api/v1/profile/addresses [post]
// @Security Bearer
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

// Address
// @Summary      Get List Address
// @Description  list address
// @Tags         profile
// @Accept       json
// @Produce      json
// @Success      200  {object}  utils.SuccessResponse
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /api/v1/profile/addresses/me [get]
// @Security Bearer
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

// Address
// @Summary      Search List Address
// @Description  search address
// @Tags         profile
// @Accept       json
// @Produce      json
// @Param        pageSize    query     string  false  "pageSize"
// @Param        pageNumber    query     string  false  "pageNumber"
// @Param        city    query     string  false  "city"
// @Param        province    query     string  false  "province"
// @Param        postalCode    query     string  false  "postalCode"
// @Param        country    query     string  false  "country"
// @Param        accountsID    query     string  false  "accountsID"
// @Param        orderBy    query     string  false  "column name"
// @Param        orderType    query     string  false  "e.g desc or asc"
// @Success      200  {object}  utils.SuccessResponse
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /api/v1/profile/addresses [get]
// @Security Bearer
func (s *ServerHttpHandler) SearchAddresses(c echo.Context) (err error) {

	// pare json
	body := new(request.SearchAddressesRequest)
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

	arg := entities.AddressesQueryParams{
		PageNumber:    body.PageNumber,
		PageSize:      body.PageSize,
		City:          body.City,
		StateProvince: body.Province,
		PostalCode:    body.PostalCode,
		Country:       body.Country,
		AccountsID:    body.AccountsID,
		OrderBy:       body.OrderBy,
		OrderType:     body.OrderType,
	}

	result, err := s.AddressesUsecase.SearchAddresses(arg)
	if err != nil {
		return utils.RespondWithError(http.StatusInternalServerError, err.Error())
	}

	message := "get addresses completed"
	return utils.RespondWithJSON(c, http.StatusOK, message, result)
}

// Address
// @Summary      Update Address
// @Description  update address
// @Tags         profile
// @Accept       json
// @Produce      json
// @Param        address_id   path      string  true  "Address ID"
// @Param request body request.UpdateAddressesRequest true "body request"
// @Success      200  {object}  utils.SuccessResponse
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /api/v1/profile/addresses/{address_id} [put]
// @Security Bearer
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

// Delete Address
// @Summary      Delete Address By Address Id
// @Description  Delete Address
// @Tags         profile
// @Accept       json
// @Produce      json
// @Param        address_id   path      string  true  "Address ID"
// @Success      200  {object}  utils.SuccessResponse
// @Failure      400  {object}  utils.ErrorResponse
// @Failure      404  {object}  utils.ErrorResponse
// @Failure      500  {object}  utils.ErrorResponse
// @Router       /api/v1/profile/addresses/{address_id} [delete]
// @Security Bearer
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
	return utils.RespondWithJSON(c, http.StatusOK, message, payload)
}
