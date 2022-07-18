package router

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"

	"golang_blog/models"
	"golang_blog/workList"
)

// @Summary register
// @Description register
// @Tags User
// @Param data body models.User true "用户"
// @Accept json
// @Produce json
// @Success 200
// @Router /register [post]
func reg(c echo.Context) error {
	var user models.User
	if err := c.Bind(&user); err != nil {
		Render(c, err)
		return err
	}
	err := workList.NewWorkList().Reg(c, &user)
	if err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil)
	return nil
}

// @Summary login
// @Description login
// @Tags User
// @Param data body models.User true "用户"
// @Accept json
// @Produce json
// @Success 200
// @Router /login [post]
func login(c echo.Context) error {
	// 数据绑定
	var user models.User
	if err := c.Bind(&user); err != nil {
		return err
	}
	token, err := workList.NewWorkList().Login(c, &user)
	if err != nil {
		Render(c, err)
		return err
	}
	cookie := &http.Cookie{
		Name:    "Auth",
		Value:   token.Auth,
		Expires: time.Now().Add(5 * time.Minute),
	}
	c.SetCookie(cookie)
	Render(c, err, token)
	return nil
}

// @Summary update_user
// @Description update user
// @Tags User
// @Security ApiKeyAuth
// @Param data body models.User true "用户"
// @Accept json
// @Produce json
// @Success 200
// @Router /user/{id} [put]
func updateUser(c echo.Context) error {
	idString := c.Param("user_id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		Render(c, err)
		return err
	}
	var user = new(models.User)
	if err = c.Bind(&user); err != nil {
		Render(c, err)
		return err
	}
	if err = workList.NewWorkList().UpdateUser(c, int32(id), user); err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil)
	return nil
}

// @Summary delete_user
// @Description delete user
// @Tags User
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Success 200
// @Router /user/{id} [delete]
func deleteUser(c echo.Context) error {
	idString := c.Param("user_id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		Render(c, err)
		return err
	}
	var userParam = new(models.User)
	if err = c.Bind(&userParam); err != nil {
		Render(c, err)
		return err
	}
	if err = workList.NewWorkList().DeleteUser(c, int32(id)); err != nil {
		Render(c, err)
		return err
	}
	Render(c, nil)
	return nil
}
