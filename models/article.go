package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Article struct {
	ID             int32  `gorm:"primary_key"`
	ContentID      int32  `gorm:"column:content_id"` // 关联到 content 表
	UserID         int32  `gorm:"column:user_id"`    // 关联到 user 表
	ArticleName    string `gorm:"column:article_name" json:"article_name"`
	ArticleContent string `gorm:"column:article_content" json:"article_content"`
	CreateTime     int64  `gorm:"column:create_time"`
	UpdateTime     int64  `gorm:"column:update_time"`
}

// 声明对象的时候直接返回一个接口
func NewArticle() Moder {
	return &Article{}
}

func (t *Article) TableName() string {
	return "articles"
}

func (t *Article) Insert(data interface{}) (err error) {
	t.CreateTime = time.Now().Unix()
	t.UpdateTime = time.Now().Unix()
	return DB.Table(t.TableName()).Create(&t).Error
}

func (t *Article) Delete(query interface{}, args ...interface{}) (err error) {
	return DB.Table(t.TableName()).Where(query, args...).Delete(&t).Error
}

func (t *Article) Save(data interface{}, query interface{}, args ...interface{}) (err error) {
	return DB.Table(t.TableName()).Omit("id").Where(query, args...).Save(&t).Error
}

func (t *Article) WhereOne(query interface{}, args ...interface{}) (article interface{}, err error) {
	err = DB.Table(t.TableName()).Where(query, args...).First(&article).Error
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (t *Article) WhereAll(query interface{}, args ...interface{}) (articles interface{}, err error) {
	err = DB.Table(t.TableName()).Where(query, args...).Find(&articles).Error
	if err != nil {
		return nil, err
	}
	return articles, nil
}

func (t *Article) PageList(params *ListPageInput, conditions interface{}, args ...interface{}) (articles interface{}, count int64, err error) {
	offset := (params.Page - 1) * params.PageSize
	query := DB.Table(t.TableName()).Where(conditions, args...)

	err = query.Offset(offset).Limit(params.PageSize).Find(&articles).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}

	err = DB.Table(t.TableName()).Where(conditions, args...).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}
	return articles, count, nil
}
