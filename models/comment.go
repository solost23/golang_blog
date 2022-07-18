package models

import (
	"errors"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID         int32  `gorm:"column:user_id" json:"user_id"`
	ArticleID      int32  `gorm:"column:article_id" json:"article_id"`
	ParentID       int32  `gorm:"column:parent_id" json:"parent_id"`
	CommentContent string `gorm:"column:comment_content" json:"comment_content"`
	IsThumbsUp     string `gorm:"column:is_thumbs_up;type:enum('COMMENT','THUMBSUP');default:'THUMBSUP'" json:"is_thumbs_up"`
}

func NewComment() Moder {
	return &Comment{}
}

func (t *Comment) TableName() string {
	return "comments"
}

func (t *Comment) Insert() (err error) {
	return DB.Table(t.TableName()).Create(&t).Error
}

func (t *Comment) Delete(query interface{}, args ...interface{}) (err error) {
	return DB.Table(t.TableName()).Where(query, args...).Delete(&t).Error
}

func (t *Comment) Save(query interface{}, args ...interface{}) (err error) {
	return DB.Table(t.TableName()).Where(query, args...).Save(&t).Error
}

func (t *Comment) WhereOne(query interface{}, args ...interface{}) (interface{}, error) {
	var comment = new(Comment)
	var err error
	err = DB.Table(t.TableName()).Where(query, args...).First(comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (t *Comment) WhereAll(query interface{}, args ...interface{}) (interface{}, error) {
	var comments []*Comment
	var err error
	err = DB.Table(t.TableName()).Where(query, args...).Find(comments).Error
	if err != nil {
		return comments, err
	}
	return comments, nil
}

func (t *Comment) PageList(params *ListPageInput, conditions interface{}, args ...interface{}) (comments interface{}, count int64, err error) {
	offset := (params.Page - 1) * params.PageSize
	query := DB.Table(t.TableName()).Where(conditions, args...)

	err = query.Offset(offset).Limit(params.PageSize).Find(comments).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return comments, 0, err
	}

	err = DB.Table(t.TableName()).Where(conditions, args...).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return comments, 0, err
	}
	return comments, count, nil
}
