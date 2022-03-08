package workList

import (
	"errors"
	"fmt"
	"log"

	"golang_blog/model"
)

func (w *WorkList) CreateContent(content *model.Content) error {
	// 通过用户名字获取用户id
	var user model.User
	user.UserName = w.ctx.Get("token").(string)
	fmt.Println(user)
	if err := user.FindByName(); err != nil {
		return err
	}
	// 将用户id给到content.user_id
	content.UserID = user.ID
	// 查询此分类是否存在，如存在，则返回错误
	if err := content.FindByUserIdAndContentName(); err == nil {
		return errors.New("此用户下此分类已存在，不可重复创建")
	}
	// 若存在，则创建分类
	if err := content.Create(); err != nil {
		return nil
	}
	return nil
}

func (w *WorkList) DeleteContent(content *model.Content) error {
	// 通过名字获取用户id
	var user model.User
	user.UserName = w.ctx.Get("token").(string)
	if err := user.FindByName(); err != nil {
		return err
	}
	// 将用户id给到content.user_id
	content.UserID = user.ID
	// 查询此分类是否存在
	// 若不存在，则返回错误
	if err := content.FindByUserIdAndContentName(); err != nil {
		return err
	}
	//若存在，则删除
	if err := content.Delete(); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) UpdateContent(content *model.Content) error {
	var user model.User
	user.UserName = w.ctx.Get("token").(string)
	if err := user.FindByName(); err != nil {
		return err
	}
	var tmpContent model.Content
	tmpContent.UserID = user.ID
	tmpContent.ContentName = content.ContentName
	if err := tmpContent.FindByUserIdAndContentName(); err != nil {
		return err
	}

	content.ID = tmpContent.ID
	if err := content.Update(); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) GetAllContent(content *model.Content) ([]*model.Content, error) {
	// 直接查询就可以
	log.Println("in content DB")
	contentList, err := content.Find()
	if err != nil {
		return contentList, err
	}
	return contentList, err
}
