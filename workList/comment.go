package workList

import (
	"errors"
	"strconv"
	"strings"

	"gorm.io/gorm"

	"golang_blog/models"
)

func (w *WorkList) GetAllComment(commentParam *models.Comment) ([]*models.Comment, error) {
	articleID := w.ctx.Get("article_id").(string)
	// 先查一遍articleID是否有这篇文章，如果没有则直接返回错误
	// 如果有，则返回此篇文章的所有评论
	articleIdInt, err := strconv.Atoi(articleID)
	if err != nil {
		return nil, err
	}

	query := []string{"id = ?"}
	args := []interface{}{articleIdInt}
	_, err = models.NewArticle().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	query = []string{"article_id = ?", "is_thumbs_up = ?"}
	args = []interface{}{commentParam.ArticleID, commentParam.IsThumbsUp}
	comments, err := models.NewComment().WhereAll(strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return comments.([]*models.Comment), nil
}

func (w *WorkList) CreateComment(commentParam *models.Comment) error {
	// base logic: 查询文章是否存在，如果不存在，则参数错误
	// 创建
	query := []string{"id = ?"}
	args := []interface{}{commentParam.ArticleID}
	_, err := models.NewArticle().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = models.NewComment().Insert(&commentParam)
	if err != nil {
		return err
	}
	return nil
}

func (w *WorkList) DeleteComment(commentParam *models.Comment) error {
	// 查询评论是否存在，存在则删除
	query := []string{"id = ?"}
	args := []interface{}{commentParam.ID}
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
