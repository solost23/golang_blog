package models

import (
	"errors"

	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	CategoryID     int32  `gorm:"column:category_id"` // 关联到 categories 表
	UserID         int32  `gorm:"column:user_id"`     // 关联到 users 表
	ArticleName    string `gorm:"column:article_name" json:"article_name"`
	ArticleContent string `gorm:"column:article_content" json:"article_content"`
}

// 声明对象的时候直接返回一个接口
func NewArticle() Moder {
	return &Article{}
}

func (t *Article) TableName() string {
	return "articles"
}

func (t *Article) Insert() (err error) {
	return DB.Table(t.TableName()).Create(&t).Error
}

func (t *Article) Delete(query interface{}, args ...interface{}) (err error) {
	return DB.Table(t.TableName()).Where(query, args...).Delete(&t).Error
}

func (t *Article) Save(query interface{}, args ...interface{}) (err error) {
	return DB.Table(t.TableName()).Omit("id").Where(query, args...).Save(&t).Error
}

func (t *Article) WhereOne(query interface{}, args ...interface{}) (interface{}, error) {
	var article = new(Article)
	err := DB.Table(t.TableName()).Where(query, args...).First(article).Error
	if err != nil {
		return article, err
	}
	return article, nil
}

func (t *Article) WhereAll(query interface{}, args ...interface{}) (interface{}, error) {
	var articles []*Article
	var err error
	err = DB.Table(t.TableName()).Where(query, args...).Find(&articles).Error
	if err != nil {
		return articles, err
	}
	return articles, nil
}

func (t *Article) PageList(params *ListPageInput, conditions interface{}, args ...interface{}) (interface{}, int64, error) {
	var articles []*Article
	var count int64
	var err error
	offset := (params.Page - 1) * params.PageSize
	query := DB.Table(t.TableName()).Where(conditions, args...)

	err = query.Offset(offset).Limit(params.PageSize).Find(articles).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return articles, 0, err
	}

	err = DB.Table(t.TableName()).Where(conditions, args...).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return articles, 0, err
	}
	return articles, count, nil
}
