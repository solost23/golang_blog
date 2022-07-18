package workList

import (
	"errors"
	"strings"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"

	"golang_blog/models"
)

func (w *WorkList) GetAllComment(_ echo.Context, articleId int32) ([]*models.Comment, error) {
	// 返回此篇文章的所有评论
	query := []string{"article_id = ?", "is_thumbs_up = ?"}
	args := []interface{}{articleId, COMMENT}
	comments, err := models.NewComment().WhereAll(strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return comments.([]*models.Comment), nil
}

func (w *WorkList) CreateComment(_ echo.Context, commentParam *models.Comment) error {
	// base logic: 查询文章是否存在，如果不存在，则参数错误
	// 创建
	query := []string{"id = ?"}
	args := []interface{}{commentParam.ArticleID}
	_, err := models.NewArticle().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = commentParam.Insert()
	if err != nil {
		return err
	}
	return nil
}

func (w *WorkList) DeleteComment(_ echo.Context, id int32) error {
	// 查询评论是否存在，存在则删除
	query := []string{"id = ?"}
	args := []interface{}{id}
	_, err := models.NewComment().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = models.NewComment().Delete(strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	return nil
}
