package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID           int32  `gorm:"primary_key"`
	UserID       int32  `gorm:"column:user_id"` // 关联到 user 表
	CategoryName string `gorm:"column:content_name" json:"category_name"`
	Introduce    string `gorm:"column:introduce" json:"introduce"`
	CreateTime   int64  `gorm:"column:create_time"`
	UpdateTime   int64  `gorm:"column:update_time"`
}

func NewCategory() Moder {
	return &Category{}
}

func (t *Category) TableName() string {
	return "contents"
}

func (t *Category) Insert(data interface{}) error {
	t.CreateTime = time.Now().Unix()
	t.UpdateTime = time.Now().Unix()
	return DB.Table("content").Create(&t).Error
}

func (t *Category) Delete(query interface{}, args ...interface{}) (err error) {
	return DB.Table(t.TableName()).Where(query, args...).Delete(&t).Error
}

func (t *Category) Save(data interface{}, query interface{}, args ...interface{}) (err error) {
	return DB.Table(t.TableName()).Where(query, args...).Save(&t).Error
}

func (t *Category) WhereOne(query interface{}, args ...interface{}) (content interface{}, err error) {
	err = DB.Table(t.TableName()).Where(query, args...).First(&content).Error
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (t *Category) WhereAll(query interface{}, args ...interface{}) (contents interface{}, err error) {
	err = DB.Table(t.TableName()).Where(query, args...).Find(&contents).Error
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func (t *Category) PageList(params *ListPageInput, conditions interface{}, args ...interface{}) (contents interface{}, count int64, err error) {
	offset := (params.Page - 1) * params.PageSize
	query := DB.Table(t.TableName()).Where(conditions, args...)

	err = query.Offset(offset).Limit(params.PageSize).Find(&contents).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}

	err = DB.Table(t.TableName()).Where(conditions, args...).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}
	return contents, count, nil
}
