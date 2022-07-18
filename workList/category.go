package workList

import (
	"errors"
	"golang_blog/middleware/jwt"
	"strings"

	"github.com/labstack/echo/v4"

	"gorm.io/gorm"

	"golang_blog/models"
)

func (w *WorkList) CreateCategory(c echo.Context, categoryParam *models.Category) error {
	// base logic: 查询此用户下分类是否存在，若不存在，则不创建，否则创建
	user := jwt.GetUser(c)
	query := []string{"user_id = ?", "category_name = ?"}
	args := []interface{}{user.Id, categoryParam.CategoryName}
	_, err := models.NewCategory().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	// 创建分类信息
	err = categoryParam.Insert()
	if err != nil {
		return err
	}
	return nil
}

func (w *WorkList) DeleteCategory(_ echo.Context, id int32) error {
	// base logic: 直接查，查到就删除
	query := []string{"id = ?"}
	args := []interface{}{id}
	_, err := models.NewCategory().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	// 删除数据
	if err = models.NewCategory().Delete(strings.Join(query, " AND "), args...); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) UpdateCategory(_ echo.Context, id int32, categoryParam *models.Category) error {
	// base logic: 查询分类id 如果有，则更新数据
	query := []string{"id = ?"}
	args := []interface{}{id}
	_, err := models.NewCategory().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	// 更新数据
	if err = categoryParam.Save(strings.Join(query, " AND "), args...); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) GetAllCategory(_ echo.Context) ([]*models.Category, error) {
	// 直接查询
	query := []string{"1 = ?"}
	args := []interface{}{1}
	categories, err := models.NewCategory().WhereAll(strings.Join(query, " AND "), args...)
	if err != nil {
		return nil, err
	}
	return categories.([]*models.Category), nil

}
