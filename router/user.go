package router

import (
	"github.com/pkg/errors"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"golang_blog/model"
	"golang_blog/mysql"
	"golang_blog/workList"
)

// PingExample godoc
// @Summary ping user
// @Schemes
// @Description do ping
// @Tags User
// @Accept json
// @Produce json
// @Success 200
// @Router /user/register [post]
func reg(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, "request data err")
		return err
	}
	var DB = mysql.DB
	err := workList.NewWorkList(c, DB).Reg(&user)

	if err.Error() == errors.New("user data exist").Error() {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, "internal error")
		return err
	}
	c.JSON(http.StatusOK, user)
	return nil
}

// PingExample godoc
// @Summary ping user
// @Schemes
// @Description do ping
// @Tags User
// @Accept json
// @Produce json
// @Success 200
// @Router /user/login [post]
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
