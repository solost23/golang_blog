package model

import (
	"crypto/md5"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"

	"golang_blog/mysql"
)

type User struct {
	ID          int32  `gorm:"primary_key"`
	UserName    string `gorm:"column:user_name" json:"user_name"`
	Role        string `gorm:"column:role;type:enum('ADMIN','USER');default:USER" json:"role"`
	PassWord    string `gorm:"column:password" json:"password"`
	NickName    string `gorm:"column:nick_name" json:"nick_name"`
	MainPageUrl string `gorm:"column:main_page_url"`
	CreateTime  int64  `gorm:"column:create_time"`
	UpdateTime  int64  `gorm:"column:update_time"`
}

func (u *User) TableName() string {
	return "user"
}

var DB *gorm.DB = mysql.DB
var DBCasbin *gorm.DB = mysql.DBCasbin

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
	if err := DB.Table("user").Omit("id", "password").Where("id=?", u.ID).Save(u).Error; err != nil {
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

func (u *User) FindByName() error {
	if err := DB.Table("user").Where("user_name=?", u.UserName).First(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return err
		}
	}
	return nil
}

// 给数据库中用户密码进行加密
func NewMd5(str string, salt ...interface{}) string {
	if len(salt) > 0 {
		slice := make([]string, len(salt)+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
