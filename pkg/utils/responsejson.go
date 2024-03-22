package utils

import (
	"github.com/labstack/echo/v4"
)

type SuccessResponse struct {
	Status  string      `json:"status" example:"success"`
	Message string      `json:"message" example:"do something completed"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Status  string `json:"status" example:"error"`
	Message string `json:"message" example:"something went wrong"`
}

func RespondWithError(code int, message string) error {
	return echo.NewHTTPError(code, message)
}

func RespondWithJSON(c echo.Context, code int, message string, payload interface{}) error {
	return c.JSON(code, SuccessResponse{Message: message, Status: "success", Data: payload})
}
