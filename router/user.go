package router

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"jwt-go/model"
	"jwt-go/mysql"
	"jwt-go/workList"
)

func reg(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, "request data err")
		return err
	}
	var DB = mysql.DB
	err := workList.NewWorkList(c, DB).Reg(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal error")
		return err
	}
	c.JSON(http.StatusOK, user)
	return nil
}

func login(c echo.Context) error {
	// 数据绑定
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, "request data err")
		return err
	}
	var DB = mysql.DB
	token, err := workList.NewWorkList(c, DB).Login(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return err
	}
	cookie := &http.Cookie{
		Name:    "Auth",
		Value:   token.Auth,
		Expires: time.Now().Add(5 * time.Minute),
	}
	c.SetCookie(cookie)
	c.JSON(http.StatusOK, token)
	return nil
}
