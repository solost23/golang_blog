package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName    string `gorm:"column:user_name" json:"user_name"`
	Role        string `gorm:"column:role;type:enum('ADMIN','USER');default:USER" json:"role"`
	PassWord    string `gorm:"column:password" json:"password"`
	NickName    string `gorm:"column:nick_name" json:"nick_name"`
	MainPageUrl string `gorm:"column:main_page_url"`
}

func NewUser() Moder {
	return &User{}
}

func (t *User) TableName() string {
	return "users"
}

func (t *User) Insert() (err error) {
	return DB.Table(t.TableName()).Create(&t).Error
}

func (t *User) Delete(query interface{}, args ...interface{}) (err error) {
	return DB.Table(t.TableName()).Where(query, args...).Delete(&t).Error
}

func (t *User) Save(query interface{}, args ...interface{}) (err error) {
	return DB.Table(t.TableName()).Where(query, args...).Save(&t).Error
}

func (t *User) WhereOne(query interface{}, args ...interface{}) (interface{}, error) {
	var user = new(User)
	err := DB.Table(t.TableName()).Where(query, args...).First(user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (t *User) WhereAll(query interface{}, args ...interface{}) (users interface{}, err error) {
	err = DB.Table(t.TableName()).Where(query, args...).Find(users).Error
	if err != nil {
		return users, err
	}
	return users, nil
}

func (t *User) PageList(params *ListPageInput, conditions interface{}, args ...interface{}) (users interface{}, count int64, err error) {
	offset := (params.Page - 1) * params.PageSize
	query := DB.Table(t.TableName()).Where(conditions, args...)

	err = query.Offset(offset).Limit(params.PageSize).Find(users).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return users, 0, err
	}

	err = DB.Table(t.TableName()).Where(conditions, args...).Count(&count).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return users, 0, err
	}
	return users, count, nil
}
