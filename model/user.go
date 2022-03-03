package model

import (
	"time"

	"github.com/jinzhu/gorm"

	"jwt-go/mysql"
)

type User struct {
	ID         int32  `gorm:"primary_key"`
	UserName   string `gorm:"column:user_name" json:"user_name"`
	PassWord   string `gorm:"column:password" json:"password"`
	NickName   string `gorm:"column:nick_name" json:"nick_name"`
	CreateTime int64  `gorm:"column:create_time"`
	UpdateTime int64  `gorm:"column:update_time"`
}

func (u *User) TableName() string {
	return "user"
}

var DB *gorm.DB = mysql.DB

func (u *User) Create() error {
	u.CreateTime = time.Now().Unix()
	u.UpdateTime = time.Now().Unix()
	if err := DB.Table("user").Create(u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) Delete() error {
	if err := DB.Table("user").Where("id=?", u.ID).Delete(u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) Update() error {
	u.UpdateTime = time.Now().Unix()
	if err := DB.Table("user").Where("id=?", u.ID).Update(u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) FindById() error {
	if err := DB.Table("user").Where("id=?", u.ID).First(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
	}
	return nil
}

func (u *User) FindByName(userName string) error {
	if err := DB.Table("user").Where("user_name=?", userName).First(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
	}
	return nil
}
