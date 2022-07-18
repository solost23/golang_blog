package models

import (
	"errors"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	UserID       int32  `gorm:"column:user_id"` // 关联到 user 表
	CategoryName string `gorm:"column:content_name" json:"category_name"`
	Introduce    string `gorm:"column:introduce" json:"introduce"`
}

func NewCategory() Moder {
	return &Category{}
}

func (t *Category) TableName() string {
	return "categories"
}

func (t *Category) Insert() error {
	return DB.Table(t.TableName()).Create(&t).Error
}

func (t *Category) Delete(query interface{}, args ...interface{}) (err error) {
	return DB.Table(t.TableName()).Where(query, args...).Delete(&t).Error
}

func (t *Category) Save(query interface{}, args ...interface{}) (err error) {
	return DB.Table(t.TableName()).Where(query, args...).Save(&t).Error
}

func (t *Category) WhereOne(query interface{}, args ...interface{}) (interface{}, error) {
	var category = new(Category)
	var err error
	err = DB.Table(t.TableName()).Where(query, args...).First(category).Error
	if err != nil {
		return category, err
	}
	return category, nil
}

func (t *Category) WhereAll(query interface{}, args ...interface{}) (interface{}, error) {
	var categories []*Category
	var err error
	err = DB.Table(t.TableName()).Where(query, args...).Find(categories).Error
	if err != nil {
		return categories, err
	}
	return categories, nil
}

func (t *Category) PageList(params *ListPageInput, conditions interface{}, args ...interface{}) (interface{}, int64, error) {
	var categories []*Category
	var count int64
	var err error
	offset := (params.Page - 1) * params.PageSize
	query := DB.Table(t.TableName()).Where(conditions, args...)

	err = query.Offset(offset).Limit(params.PageSize).Find(categories).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return categories, 0, err
	}

	err = DB.Table(t.TableName()).Where(conditions, args...).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return categories, 0, err
	}
	return categories, count, nil
}
