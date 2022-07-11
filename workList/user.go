package workList

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"

	"golang_blog/middleware/jwt"
	"golang_blog/models"
)

func (w *WorkList) Reg(userParam *models.User) error {
	// 查看用户是否存在
	// 若存在，则返回错误
	query := []string{"user_name = ?"}
	args := []interface{}{userParam.UserName}
	user, err := models.NewUser().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("user data exist")
	}
	// 若不存在，则创建
	userParam.PassWord = models.NewMd5(userParam.PassWord, "ty")
	if err = models.NewUser().Insert(user); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) Login(userParam *models.User) (models.Token, error) {
	// 查看有无用户
	// 如果没有直接报错
	query := []string{"user_name = ?", "role = ?", "password = ?"}
	args := []interface{}{userParam.UserName, userParam.Role, models.NewMd5(userParam.PassWord, "ty")}

	var token models.Token
	_, err := models.NewUser().WhereOne(query, args...)
	if err != nil {
		return token, err
	}

	// 否则生成一个 token
	tokenString, err := jwt.CreateToken(userParam.UserName, userParam.Role)
	if err != nil {
		fmt.Println(err.Error())
		return token, err
	}
	token = models.Token{
		Auth: tokenString,
	}
	return token, nil
}

func (w *WorkList) UpdateUser(userParam *models.User) error {
	// 通过用户名查出来用户id
	query := []string{"user_name = ?"}
	args := []interface{}{userParam.UserName}
	user, err := models.NewUser().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}

	// 通过用户id更新用户信息
	query = []string{"id = ?"}
	args = []interface{}{user.(*models.User).ID}
	userParam.PassWord = models.NewMd5(userParam.PassWord, "ty")
	if err = models.NewUser().Save(userParam, query, args...); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) DeleteUser(userParam *models.User) error {
	// 查找此用户id是否存在，存在则删除用户
	query := []string{"id = ?"}
	args := []interface{}{userParam.ID}
	err := models.NewUser().Delete(strings.Join(query, " AND"), args...)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	return nil
}
