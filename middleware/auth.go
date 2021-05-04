package middleware

import (
	"strings"

	"github.com/labstack/echo"
)

type AuthContext struct {
	echo.Context
	token string
}

// BearerAuth middleware extract the authorization bearer header into a struct that can be used
// later on as needed.
func BearerAuth() echo.MiddlewareFunc {
	return func(hf echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorizationHeader := c.Request().Header.Get("Authorization")
			splits := strings.Split(authorizationHeader, " ")
			if len(splits) != 2 {
				return hf(c)
			}

			if strings.ToLower(splits[0]) != "bearer" {
				return hf(c)
			}

			return hf(&AuthContext{c, splits[1]})
		}
	}
}
