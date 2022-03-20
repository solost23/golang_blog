package jwt

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"golang_blog/common"
)

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("Auth")
		if err != nil {
			c.JSON(http.StatusBadRequest, "Cookie err")
			return err
		}
		tknStr := cookie.Value
		claims := &common.Claims{}
		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte("my_secret_key"), nil
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, "internal err")
			return err
		}
		if !tkn.Valid {
			c.JSON(http.StatusInternalServerError, "token err valid")
			return err
		}
		c.Set("token", claims.UserName)
		c.Set("claims", claims)
		next(c)
		return nil
	}
}
