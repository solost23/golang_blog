package workList

import (
	"errors"
	"fmt"

	"golang_blog/models"
)

func (w *WorkList) CreateArticle(article *models.Article) error {
	// 获取用户名
	userName := w.ctx.Get("token").(string)
	contentName := w.ctx.Get("content_name").(string)
	// 通过用户名获取用户id
	var user models.User
	user.UserName = userName
	if err := user.FindByName(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	// 获取分类名
	// 通过用户id与分类名查询有无此分类
	var content models.Content
	content.UserID = user.ID
	content.ContentName = contentName
	if err := content.FindByUserIdAndContentName(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	// 查询有无此文章
	article.ContentID = content.ID
	article.UserID = user.ID
	if err := article.FindByUserIdAndContentIdAndArticleName(); err == nil {
		return errors.New("此文章已存在，不可重复创建")
	}
	// 若没有，则插入文章
	if err := article.Create(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (w *WorkList) DeleteArticle(article *models.Article) error {
	// 获取用户名和分类名
	userName := w.ctx.Get("token").(string)
	contentName := w.ctx.Get("content_name").(string)
	articleName := w.ctx.Get("article_name").(string)
	// 查询用户id
	var user models.User
	user.UserName = userName
	if err := user.FindByName(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	// 查询分类是否存在
	var content models.Content
	content.UserID = user.ID
	content.ContentName = contentName
	if err := content.FindByNameAndUserId(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	// 从文章表中查询文章是否存在
	article.UserID = user.ID
	article.ContentID = content.ID
	article.ArticleName = articleName
	if err := article.FindByUserIdAndContentIdAndArticleName(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	// 通过文章id删除文章
	if err := article.Delete(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (w *WorkList) UpdateArticle(article *models.Article) error {
	userName := w.ctx.Get("token").(string)
	contentName := w.ctx.Get("content_name").(string)
	articleName := w.ctx.Get("article_name").(string)
	var user models.User
	user.UserName = userName
	if err := user.FindByName(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	var content models.Content
	content.UserID = user.ID
	content.ContentName = contentName
	if err := content.FindByNameAndUserId(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	var tmpArticle models.Article
	tmpArticle.UserID = user.ID
	tmpArticle.ContentID = content.ID
	tmpArticle.ArticleName = articleName
	if err := tmpArticle.FindByUserIdAndContentIdAndArticleName(); err != nil {
		fmt.Println(err.Error())
		return err
	}

	article.ID = tmpArticle.ID
	article.UserID = user.ID
	article.ContentID = content.ID
	if err := article.Update(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (w *WorkList) GetAllArticle(article *models.Article) ([]*models.Article, error) {
	// 直接查询
	articleList, err := article.Find()
	if err != nil {
		fmt.Println(err.Error())
		return articleList, err
	}
	return articleList, nil
}

func (w *WorkList) GetArticle(article *models.Article) error {
	userName := w.ctx.Get("user_name").(string)
	contentName := w.ctx.Get("content_name").(string)
	articleName := w.ctx.Get("article_name").(string)
	// 通过用户名查找到userid
	var user models.User
	user.UserName = userName
	if err := user.FindByName(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	// 通过contentName And Userid 查找到content.id
	var content models.Content
	content.UserID = user.ID
	content.ContentName = contentName
	if err := content.FindByNameAndUserId(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	article.UserID = user.ID
	article.ContentID = content.ID
	article.ArticleName = articleName
	// 通过userid 和分类名字和文章名字去查询文章内容
	if err := article.FindByUserIdAndContentIdAndArticleName(); err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}
