package workList

import (
	"errors"
	"strings"

	"gorm.io/gorm"

	"golang_blog/models"
)

func (w *WorkList) CreateCategory(categoryParam *models.Category) error {
	// 通过用户名字获取用户id
	userName := w.ctx.Get("token").(string)
	query := []string{"user_name = ?"}
	args := []interface{}{userName}
	user, err := models.NewUser().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	query = []string{"user_id = ?", "category_name = ?"}
	args = []interface{}{user.(*models.User).ID, categoryParam.CategoryName}
	_, err = models.NewCategory().WhereOne(strings.Join(query, " AND "), args...)
	if err == nil || !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("此用户下此分类已存在，不可重复创建")
	}
	// 若不遵在，则创建分类
	categoryParam.UserID = user.(*models.User).ID
	if err = models.NewCategory().Insert(categoryParam); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) DeleteCategory(categoryParam *models.Category) error {
	userName := w.ctx.Get("token").(string)
	query := []string{"user_name = ?"}
	args := []interface{}{userName}
	user, err := models.NewUser().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	query = []string{"user_id = ?", "category_name = ?"}
	args = []interface{}{user.(*models.Category).ID, categoryParam.CategoryName}
	category, err := models.NewCategory().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	query = []string{"id = ?"}
	args = []interface{}{category.(*models.Category).ID}
	if err = models.NewCategory().Delete(strings.Join(query, " AND "), args...); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) UpdateCategory(categoryParam *models.Category) error {
	userName := w.ctx.Get("token").(string)
	query := []string{"user_name = ?"}
	args := []interface{}{userName}
	user, err := models.NewUser().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil {
		return err
	}
	query = []string{"user_id = ?", "category_name = ?"}
	args = []interface{}{user.(*models.Category).ID, categoryParam.CategoryName}
	category, err := models.NewCategory().WhereOne(strings.Join(query, " AND "), args...)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	// 更新数据
	query = []string{"id = ?"}
	args = []interface{}{category.(*models.Category).ID}
	if err = models.NewCategory().Save(categoryParam, strings.Join(query, " AND "), args...); err != nil {
		return err
	}
	return nil
}

func (w *WorkList) GetAllCategory(categoryParam *models.Category) ([]*models.Category, error) {
	// 直接查询
	query := []string{"1 = ?"}
	args := []interface{}{1}
	categories, err := models.NewCategory().WhereAll(strings.Join(query, " AND "), args...)
	if err != nil {
		return nil, err
	}
	return categories.([]*models.Category), nil

}
