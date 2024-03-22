package middleware

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

type jwtCustomClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}

type CreateTokenDTO struct {
	UserID    string
	ExpiresAt time.Time
	Secret    string
}

func ConfigJWT(secret string) echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte(secret),
		SuccessHandler: func(c echo.Context) {
			c.Set("accountsID", decodeToken(c))
		},
	}

	return echojwt.WithConfig(config)
}

func CreateToken(param CreateTokenDTO) (token string, err error) {
	claims := &jwtCustomClaims{
		ID: param.UserID,
	}
	claims.ExpiresAt = jwt.NewNumericDate(param.ExpiresAt)

	// Create token with claims
	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	token, err = tokenClaim.SignedString([]byte(param.Secret))
	if err != nil {
		return
	}

	return
}

func decodeToken(c echo.Context) (id string) {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	id = claims.ID

	return
}
