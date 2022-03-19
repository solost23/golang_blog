package model

import "time"

type Comment struct {
	ID             int32  `gorm:"primary_key"`
	UserID         int32  `gorm:"column:user_id"`
	ArticleID      int32  `gorm:"column:article_id"`
	ParentID       int32  `gorm:"column:parent_id"`
	CommentContent string `gorm:"column:comment_content" json:"comment_content"`
	IsThumbsUp     string `gorm:"column:is_thumbs_up;type:enum('COMMENT','THUMBSUP');default:'THUMBSUP'" json:"is_thumbs_up"`
	CreateTime     int64  `gorm:"column:create_time"`
	UpdateTime     int64  `gorm:"column:update_time"`
}

func (c *Comment) TableName() string {
	return "comment"
}

func (c *Comment) Create() error {
	c.CreateTime = time.Now().Unix()
	c.UpdateTime = time.Now().Unix()
	if err := DB.Table("comment").Create(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Comment) Delete() error {
	if err := DB.Table("comment").Where("id=?", c.ID).Delete(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Comment) Find() ([]*Comment, error) {
	var res []*Comment
	if err := DB.Table("comment").Where("article_id=?", c.ArticleID).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (c *Comment) FindByIDAndUserID() error {
	if err := DB.Table("comment").Where("id=? AND user_id=?", c.ID, c.UserID).First(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Comment) FindByID() error {
	if err := DB.Table("comment").Where("id=?", c.ParentID).First(c).Error; err != nil {
		return err
	}
	return nil
}
