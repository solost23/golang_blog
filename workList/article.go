package workList

import (
	"errors"
	"golang_blog/middleware/jwt"
	"strings"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"

	"golang_blog/models"
)

func (w *WorkList) CreateArticle(c echo.Context, articleParam *models.Article) error {
	// base logic: 查询用户 && 查询分类是否存在, 如果都存在，那么新建
	user := jwt.GetUser(c)
	query := []string{"user_id = ?", "category_id = ?"}
	args := []interface{}{user.Id, articleParam.CategoryID}
	_, err := models.NewArticle().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = articleParam.Insert()
	if err != nil {
		return err
	}
	return nil
}

func (w *WorkList) DeleteArticle(_ echo.Context, id int32) error {
	// base login: 查询文章是否存在，不存在则报错
	// 存在则删除
	query := []string{"id = ?"}
	args := []interface{}{id}
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

func (w *WorkList) UpdateArticle(_ echo.Context, id int32, articleParam *models.Article) error {
	// 查看有无此文章，有则更新
	query := []string{"id = ?"}
	args := []interface{}{id}
	_, err := models.NewArticle().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	err = articleParam.Save(strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	return nil
}

func (w *WorkList) GetAllArticle(_ echo.Context) ([]*models.Article, error) {
	// 直接查询
	query := []string{"1 = ?"}
	args := []interface{}{1}
	articleList, err := models.NewArticle().WhereAll(strings.Join(query, " AND "), args...)
	if err != nil {
		return nil, err
	}
	return articleList.([]*models.Article), nil
}

func (w *WorkList) GetArticle(_ echo.Context, id int32) (*models.Article, error) {
	// 直接查询并返回数据
	query := []string{"id = ?"}
	args := []interface{}{id}
	article, err := models.NewArticle().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	return article.(*models.Article), nil
}
