package workList

import (
	"errors"
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"

	"golang_blog/middleware/jwt"
	"golang_blog/models"
)

func (w *WorkList) Reg(_ echo.Context, userParam *models.User) error {
	// 查看用户是否存在
	// 若存在，则返回错误
	query := []string{"user_name = ?"}
	args := []interface{}{userParam.UserName}
	user, err := models.NewUser().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user data exist")
	}
	if user.(*models.User).ID != 0 {
		return errors.New("user data exist")
	}
	// 若不存在，则创建
	userParam.PassWord = models.NewMd5(userParam.PassWord, "ty")
	if err = userParam.Insert(); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) Login(_ echo.Context, userParam *models.User) (models.Token, error) {
	// 查看有无用户
	// 如果没有直接报错
	query := []string{"user_name = ?", "role = ?", "password = ?"}
	args := []interface{}{userParam.UserName, userParam.Role, models.NewMd5(userParam.PassWord, "ty")}

	var token models.Token
	user, err := models.NewUser().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil {
		return token, err
	}

	// 否则生成一个 token
	tokenString, err := jwt.CreateToken(int32(user.(*models.User).ID), user.(*models.User).UserName, user.(*models.User).Role)
	if err != nil {
		fmt.Println(err.Error())
		return token, err
	}
	token = models.Token{
		Auth: tokenString,
	}
	return token, nil
}

func (w *WorkList) UpdateUser(_ echo.Context, id int32, userParam *models.User) error {
	// 通过用户id查出来用户id
	query := []string{"id = ?"}
	args := []interface{}{id}
	user, err := models.NewUser().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}

	// 通过用户id更新用户信息
	query = []string{"id = ?"}
	args = []interface{}{user.(*models.User).ID}
	userParam.PassWord = models.NewMd5(userParam.PassWord, "ty")
	if err = userParam.Save(query, args...); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) DeleteUser(_ echo.Context, id int32) error {
	// 查找此用户id是否存在，存在则删除用户
	query := []string{"id = ?"}
	args := []interface{}{id}
	err := models.NewUser().Delete(strings.Join(query, " AND"), args...)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
