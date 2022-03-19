package model

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Content struct {
	ID          int32  `gorm:"primary_key"`
	UserID      int32  `gorm:"column:user_id"` // 关联到 user 表
	ContentName string `gorm:"column:content_name" json:"content_name"`
	Introduce   string `gorm:"column:introduce" json:"introduce"`
	CreateTime  int64  `gorm:"column:create_time"`
	UpdateTime  int64  `gorm:"column:update_time"`
}

func (c *Content) TableName() string {
	return "content"
}

func (c *Content) Create() error {
	c.CreateTime = time.Now().Unix()
	c.UpdateTime = time.Now().Unix()
	if err := DB.Table("content").Create(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Content) Delete() error {
	if err := DB.Table("content").Where("id=?", c.ID).Delete(c).Error; err != nil {
		return err
	}
	return nil
}

func (c *Content) Update() error {
	c.UpdateTime = time.Now().Unix()
	if err := DB.Table("content").Omit("id", "user_id").Where("id=?", c.ID).Save(c).Error; err != nil {
		return err
	}
	return nil
}

// 查询所有分类
func (c *Content) Find() ([]*Content, error) {
	var res []*Content
	if err := DB.Table("content").Find(&res).Error; err != nil {
		log.Fatalln(err)
		return res, err
	}
	return res, nil
}

// 根据分类名查分类id
func (c *Content) FindByNameAndUserId() error {
	if err := DB.Table("content").Where("content_name=? AND user_id=?", c.ContentName, c.UserID).First(c).Error; err != nil {
		return err
	}
	return nil
}

// 根据id查分类名称
func (c *Content) FindById() error {
	if err := DB.Table("content").Where("id=?", c.ID).First(c).Error; err != nil {
		return err
	}
	return nil
}

// 根据用户id和分类名查询记录
func (c *Content) FindByUserIdAndContentName() error {
	if err := DB.Table("content").Where("user_id=? AND content_name=?", c.UserID, c.ContentName).First(c).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
	}
	return nil
}
