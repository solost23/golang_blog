package workList

import (
	"errors"
	"golang_blog/common"
	"golang_blog/model"
)

func (w *WorkList) Reg(user *model.User) error {
	// 查看用户是否存在
	// 若存在，则返回错误
	if err := user.FindByName(); err == nil {
		return errors.New("user data exist")
	}
	// 若不存在，则创建
	user.PassWord = common.NewMd5(user.PassWord, "ty")
	if err := user.Create(); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) Login(user *model.User) (model.Token, error) {
	// 查看有无用户
	// 如果没有直接报错
	userName := user.UserName
	password := user.PassWord
	var token model.Token
	if err := user.FindByName(); err != nil {
		return token, err
	}
	// 如果有则检查账户名密码
	// 如果账户名 || 密码错误，返回错误
	if userName != user.UserName || common.NewMd5(password, "ty") != user.PassWord {
		return token, errors.New("username or password failed")
	}
	// 否则生成一个 token
	tokenString, err := common.CreateToken(userName)
	if err != nil {
		return token, err
	}
	token = model.Token{
		Auth: tokenString,
	}
	return token, nil
}

func (w *WorkList) UpdateUser(user *model.User) error {
	// 通过用户名查出来用户id
	var tmpUser model.User
	tmpUser.UserName = user.UserName
	if err := tmpUser.FindByName(); err != nil {
		return err
	}
	// 通过用户id更新用户信息
	user.ID = tmpUser.ID
	if err := user.Update(); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) DeleteUser(user *model.User) error {
	userName := w.ctx.Get("user_name").(string)
	// 通过用户名获取用户ID
	user.UserName = userName
	if err := user.FindByName(); err != nil {
		return err
	}
	// 删除用户
	if err := user.Delete(); err != nil {
		return err
	}
	return nil
}
