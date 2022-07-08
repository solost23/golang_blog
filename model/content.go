package model

import (
	"errors"
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

func NewContent() Moder {
	return &Content{}
}

func (t *Content) TableName() string {
	return "contents"
}

func (t *Content) Insert(data interface{}) error {
	t.CreateTime = time.Now().Unix()
	t.UpdateTime = time.Now().Unix()
	return DB.Table("content").Create(&t).Error
}

func (t *Content) Delete(query interface{}, args ...interface{}) (err error) {
	return DB.Table(t.TableName()).Where(query, args...).Delete(&t).Error
}

func (t *Content) Save(data interface{}) (err error) {
	return DB.Table(t.TableName()).Save(&t).Error
}

func (t *Content) WhereOne(query interface{}, args ...interface{}) (content interface{}, err error) {
	err = DB.Table(t.TableName()).Where(query, args...).First(&content).Error
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (t *Content) WhereAll(query interface{}, args ...interface{}) (contents interface{}, err error) {
	err = DB.Table(t.TableName()).Where(query, args...).Find(&contents).Error
	if err != nil {
		return nil, err
	}
	return contents, nil
}

func (t *Content) PageList(params *ListPageInput, conditions interface{}, args ...interface{}) (contents interface{}, count int64, err error) {
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

//func (c *Content) Update() error {
//	c.UpdateTime = time.Now().Unix()
//	if err := DB.Table("content").Omit("id", "user_id").Where("id=?", c.ID).Save(c).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//// 查询所有分类
//func (c *Content) Find() ([]*Content, error) {
//	var res []*Content
//	if err := DB.Table("content").Find(&res).Error; err != nil {
//		log.Fatalln(err)
//		return res, err
//	}
//	return res, nil
//}
//
//// 根据分类名查分类id
//func (c *Content) FindByNameAndUserId() error {
//	if err := DB.Table("content").Where("content_name=? AND user_id=?", c.ContentName, c.UserID).First(c).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//// 根据id查分类名称
//func (c *Content) FindById() error {
//	if err := DB.Table("content").Where("id=?", c.ID).First(c).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//// 根据用户id和分类名查询记录
//func (c *Content) FindByUserIdAndContentName() error {
//	if err := DB.Table("content").Where("user_id=? AND content_name=?", c.UserID, c.ContentName).First(c).Error; err != nil {
//		if err == gorm.ErrRecordNotFound {
//			return err
//		}
//	}
//	return nil
//}
