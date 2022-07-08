package workList

import (
	"errors"
	"fmt"
	"strconv"

	"gorm.io/gorm"

	"golang_blog/models"
)

func (w *WorkList) GetAllComment(comment *models.Comment) ([]*models.Comment, error) {
	articleID := w.ctx.Get("article_id").(string)
	// 先查一遍articleID是否有这篇文章，如果没有则直接返回错误
	// 如果有，则调用返回评论按钮
	var res []*models.Comment
	var article models.Article
	articleIDInt, err := strconv.Atoi(articleID)
	if err != nil {
		fmt.Println(err.Error())
		return res, err
	}
	article.ID = int32(articleIDInt)
	if err := article.FindById(); err != nil {
		fmt.Println(err.Error())
		return res, err
	}
	comment.ArticleID = article.ID
	res, err = comment.Find()
	if err != nil {
		fmt.Println(err.Error())
		return res, err
	}
	return res, nil
}

func (w *WorkList) CreateComment(comment *models.Comment) error {
	userName := w.ctx.Get("user_name")
	articleID := w.ctx.Get("article_id")
	parentID := w.ctx.Get("parent_id")
	// 查看有无此用户，若没有，则返回错误
	// 查看有无此文章，若没有，则返回错误
	// 存储评论
	articleIDInt, err := strconv.Atoi(articleID.(string))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	parentIDInt, err := strconv.Atoi(parentID.(string))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	var user = new(models.User)
	user.UserName = userName.(string)
	if err := user.FindByName(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	var article = new(models.Article)
	article.ID = int32(articleIDInt)
	if err := article.FindById(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if comment.IsThumbsUp != "COMMENT" && comment.IsThumbsUp != "THUMBSUP" {
		return errors.New("评论类型不存在")
	}
	comment.UserID = user.ID
	comment.ArticleID = article.ID
	comment.ParentID = int32(parentIDInt)
	var tmpComment = new(models.Comment)
	tmpComment.ParentID = int32(parentIDInt)
	// 查找当前父评论是否存在，若不存在，则自己就是父评论
	if err := tmpComment.FindByID(); err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println(err.Error())
			return err
		}
		comment.ParentID = 0
	}
	if err := comment.Create(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (w *WorkList) DeleteComment(comment *models.Comment) error {
	userName := w.ctx.Get("user_name")
	commentID := w.ctx.Get("comment_id")
	// 查找当前用户是否存在，若不存在，则返回错误
	// 查找当前用户是否有此评论，如果有，则删除
	commentIDInt, err := strconv.Atoi(commentID.(string))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	var user = new(models.User)
	user.UserName = userName.(string)
	if err := user.FindByName(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	comment.ID = int32(commentIDInt)
	comment.UserID = user.ID
	if err := comment.FindByIDAndUserID(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	if err := comment.Delete(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
