package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID             int32  `gorm:"primary_key" json:"id"`
	UserID         int32  `gorm:"column:user_id" json:"user_id"`
	ArticleID      int32  `gorm:"column:article_id" json:"article_id"`
	ParentID       int32  `gorm:"column:parent_id" json:"parent_id"`
	CommentContent string `gorm:"column:comment_content" json:"comment_content"`
	IsThumbsUp     string `gorm:"column:is_thumbs_up;type:enum('COMMENT','THUMBSUP');default:'THUMBSUP'" json:"is_thumbs_up"`
	CreateTime     int64  `gorm:"column:create_time"`
	UpdateTime     int64  `gorm:"column:update_time"`
}

func NewComment() Moder {
	return &Comment{}
}

func (t *Comment) TableName() string {
	return "comments"
}

func (t *Comment) Insert(data interface{}) (err error) {
	t.CreateTime = time.Now().Unix()
	t.UpdateTime = time.Now().Unix()
	return DB.Table(t.TableName()).Create(&t).Error
}

func (t *Comment) Delete(query interface{}, args ...interface{}) (err error) {
	return DB.Table(t.TableName()).Where(query, args...).Delete(&t).Error
}

func (t *Comment) Save(data interface{}, query interface{}, args ...interface{}) (err error) {
	return DB.Table(t.TableName()).Where(query, args...).Save(&t).Error
}

func (t *Comment) WhereOne(query interface{}, args ...interface{}) (comment interface{}, err error) {
	err = DB.Table(t.TableName()).Where(query, args...).First(&comment).Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (t *Comment) WhereAll(query interface{}, args ...interface{}) (comments interface{}, err error) {
	err = DB.Table(t.TableName()).Where(query, args...).Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}

func (t *Comment) PageList(params *ListPageInput, conditions interface{}, args ...interface{}) (comments interface{}, count int64, err error) {
	offset := (params.Page - 1) * params.PageSize
	query := DB.Table(t.TableName()).Where(conditions, args...)

	err = query.Offset(offset).Limit(params.PageSize).Find(&comments).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}

	err = DB.Table(t.TableName()).Where(conditions, args...).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}
	return comments, count, nil
}
