package model

import "time"

type Article struct {
	ID             int32  `gorm:"primary_key"`
	ContentID      int32  `gorm:"column:content_id"` // 关联到 content 表
	UserID         int32  `gorm:"column:user_id"`    // 关联到 user 表
	ArticleName    string `gorm:"column:article_name" json:"article_name"`
	ArticleContent string `gorm:"column:article_content" json:"article_content"`
	CreateTime     int64  `gorm:"column:create_time"`
	UpdateTime     int64  `gorm:"column:update_time"`
}

func (a *Article) TableName() string {
	return "article"
}

func (a *Article) Create() error {
	a.CreateTime = time.Now().Unix()
	a.UpdateTime = time.Now().Unix()
	if err := DB.Table("article").Create(a).Error; err != nil {
		return err
	}
	return nil
}

func (a *Article) Delete() error {
	if err := DB.Table("article").Where("id=?", a.ID).Delete(a).Error; err != nil {
		return err
	}
	return nil
}

func (a *Article) Update() error {
	a.UpdateTime = time.Now().Unix()
	if err := DB.Table("article").Where("id=?", a.ID).Update(a).Error; err != nil {
		return err
	}
	return nil
}

func (a *Article) Find() ([]*Article, error) {
	var res []*Article
	if err := DB.Table("article").Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

// 根据文章名查文章id
func (a *Article) FindByName(articleName string) error {
	if err := DB.Table("article").Where("article_name=?", a.ArticleName).First(a).Error; err != nil {
		return err
	}
	return nil
}

// 根据id查文章名称
func (a *Article) FindById() error {
	if err := DB.Table("article").Where("id=?", a.ID).First(a).Error; err != nil {
		return err
	}
	return nil
}

func (a *Article) FindByUserIdAndContentIdAndArticleName() error {
	if err := DB.Table("article").Where("user_id=? AND content_id=? AND article_name=?", a.UserID, a.ContentID, a.ArticleName).First(a).Error; err != nil {
		return err
	}
	return nil
}
