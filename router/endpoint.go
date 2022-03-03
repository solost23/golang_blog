package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	jwt "jwt-go/middleware/jwt"
)

func RegisterUser(router *echo.Echo) {
	user := router.Group("/user")
	{
		user.POST("/register", reg)
		user.POST("/login", login)
	}
}

func Register() *echo.Echo {
	router := echo.New()
	router.Use(middleware.Logger(), middleware.Recover())

	RegisterUser(router)

	router.GET("/index", func(c echo.Context) error {
		c.JSON(http.StatusOK, "hello world")
		return nil
	}, jwt.JWTAuth)

	// 注册一下文档接口
	return router
}
