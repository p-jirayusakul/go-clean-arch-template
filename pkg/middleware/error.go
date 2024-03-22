package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/p-jirayusakul/go-clean-arch-template/pkg/utils"
)

func ErrorHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			// Handle errors here
			switch e := err.(type) {
			case *echo.HTTPError:
				return c.JSON(e.Code, utils.ErrorResponse{Message: e.Message.(string), Status: "error"})
			default:
				return c.JSON(http.StatusInternalServerError, utils.ErrorResponse{Message: "Internal Server Error", Status: "error"})
			}
		}
		return nil
	}
}
