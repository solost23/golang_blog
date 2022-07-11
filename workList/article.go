package workList

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"golang_blog/models"
)

func (w *WorkList) CreateArticle(articleParam *models.Article) error {
	// base logic: 查询用户 && 查询分类是否存在, 如果都存在，那么新建
	query := []string{"user_id = ?", "category_id = ?"}
	args := []interface{}{articleParam.UserID, articleParam.CategoryID}
	_, err := models.NewArticle().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = models.NewArticle().Insert(articleParam)
	if err != nil {
		return err
	}
	return nil
}

func (w *WorkList) DeleteArticle(articleParam *models.Article) error {
	// base login: 查询文章是否存在，不存在则报错
	// 存在则删除
	query := []string{"id = ?"}
	args := []interface{}{articleParam.ID}
	_, err := models.NewArticle().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = models.NewArticle().Delete(strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	return nil
}

func (w *WorkList) UpdateArticle(articleParam *models.Article) error {
	// 查看有无此文章，有则更新
	query := []string{"id = ?"}
	args := []interface{}{articleParam.ID}
	_, err := models.NewArticle().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = models.NewArticle().Save(articleParam, strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	return nil
}

func (w *WorkList) GetAllArticle(articleParam *models.Article) ([]*models.Article, error) {
	// 直接查询
	query := []string{"1 = ?"}
	args := []interface{}{1}
	articleList, err := models.NewArticle().WhereAll(strings.Join(query, " AND "), args...)
	if err != nil {
		return nil, err
	}
	return articleList.([]*models.Article), nil
}

func (w *WorkList) GetArticle(articleParam *models.Article) (*models.Article, error) {
	// 直接查询并返回数据
	query := []string{"id = ?"}
	args := []interface{}{articleParam.ID}
	article, err := models.NewArticle().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return article.(*models.Article), nil
}
