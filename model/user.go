package model

import (
	"errors"
	"time"

	"gorm.io/gorm"
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

func NewUser() Moder {
	return &User{}
}

func (t *User) TableName() string {
	return "users"
}

func (t *User) Insert(data interface{}) (err error) {
	t.CreateTime = time.Now().Unix()
	t.UpdateTime = time.Now().Unix()
	return DB.Table(t.TableName()).Create(&t).Error
}

func (t *User) Delete(query interface{}, args ...interface{}) (err error) {
	return DB.Table(t.TableName()).Where(query, args...).Delete(&t).Error
}

func (t *User) Save(data interface{}) (err error) {
	t.UpdateTime = time.Now().Unix()
	return DB.Table(t.TableName()).Save(&t).Error
}

func (t *User) WhereOne(query interface{}, args ...interface{}) (user interface{}, err error) {
	err = DB.Table(t.TableName()).Where(query, args...).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (t *User) WhereAll(query interface{}, args ...interface{}) (users interface{}, err error) {
	err = DB.Table(t.TableName()).Where(query, args...).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (t *User) PageList(params *ListPageInput, conditions interface{}, args ...interface{}) (users interface{}, count int64, err error) {
	offset := (params.Page - 1) * params.PageSize
	query := DB.Table(t.TableName()).Where(conditions, args...)

	err = query.Offset(offset).Limit(params.PageSize).Find(&users).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}

	err = DB.Table(t.TableName()).Where(conditions, args...).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, 0, err
	}
	return users, count, nil
}

//func (u *User) Update() error {
//	u.UpdateTime = time.Now().Unix()
//	if err := DB.Table("user").Omit("id", "password").Where("id=?", u.ID).Save(u).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//func (u *User) FindById() error {
//	if err := DB.Table("user").Where("id=?", u.ID).First(u).Error; err != nil {
//		if err == gorm.ErrRecordNotFound {
//			return err
//		}
//	}
//	return nil
//}
//
//func (u *User) FindByName() error {
//	if err := DB.Table("user").Where("user_name=?", u.UserName).First(u).Error; err != nil {
//		if err == gorm.ErrRecordNotFound {
//			return err
//		}
//	}
//	return nil
//}
