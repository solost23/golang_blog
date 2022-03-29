package router

import (
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"

	"golang_blog/model"
	"golang_blog/mysql"
	"golang_blog/workList"
)

// @Summary register
// @Description register
// @Tags User
// @Security ApiKeyAuth
// @Param data body model.User true "用户"
// @Accept json
// @Produce json
// @Success 200
// @Router /register [post]
func reg(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, "request data err")
		return err
	}
	var DB = mysql.DB
	err := workList.NewWorkList(c, DB).Reg(&user)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}
	c.JSON(http.StatusOK, user)
	return nil
}

// @Summary login
// @Description login
// @Tags User
// @Security ApiKeyAuth
// @Param data body model.User true "用户"
// @Accept json
// @Produce json
// @Success 200
// @Router /login [post]
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
		c.JSON(http.StatusInternalServerError, err.Error())
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

// @Summary update_user
// @Description update user
// @Tags User
// @Security ApiKeyAuth
// @Param data body model.User true "用户"
// @Accept json
// @Produce json
// @Success 200
// @Router /user/{user_name} [put]
func updateUser(c echo.Context) error {
	userName := c.Param("user_name")
	c.Set("user_name", userName)
	var user model.User
	if err := c.Bind(&user); err != nil {
		c.JSON(http.StatusBadRequest, "request data err")
		return err
	}
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}
	c.JSON(http.StatusOK, user)
	return nil
}

// @Summary delete_user
// @Description delete user
// @Tags User
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /user/{user_name} [delete]
func deleteUser(c echo.Context) error {
	userName := c.Param("user_name")
	c.Set("user_name", userName)
	var user model.User
	var DB = mysql.DB
	if err := workList.NewWorkList(c, DB).DeleteUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return err
	}
	c.JSON(http.StatusOK, "user delete success")
	return nil
}
