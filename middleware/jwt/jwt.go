package jwt

import (
	"net/http"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/labstack/echo/v4"
)

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		cookie, err := c.Cookie("Auth")
		if err != nil {
			c.JSON(http.StatusBadRequest, "Cookie err")
			return err
		}
		tknStr := cookie.Value
		claims := &Claims{}
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
		c.Set("user", claims)
		next(c)
		return nil
	}
}

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Id       int32
	UserName string
	Role     string
	jwt.StandardClaims
}

func CreateToken(userId int32, userName, role string) (string, error) {
	claims := &Claims{
		Id:             userId,
		UserName:       userName,
		Role:           role,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GetUser(c echo.Context) *Claims {
	return c.Get("user").(*Claims)
}
