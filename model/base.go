package model

import (
	"crypto/md5"
	"fmt"
	"golang_blog/mysql"
	"strings"

	"gorm.io/gorm"
)

var DB *gorm.DB = mysql.DB
var DBCasbin *gorm.DB = mysql.DBCasbin

// 分页结构
type ListPageInput struct {
	Page     int `comment:"当前页"`
	PageSize int `comment:"每页记录数"`
}

type Token struct {
	Auth string
}

// 给数据库中用户密码进行加密
func NewMd5(str string, salt ...interface{}) string {
	if len(salt) > 0 {
		slice := make([]string, len(salt)+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	return fmt.Sprintf("%x", md5.Sum([]byte(str)))
}
