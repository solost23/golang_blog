package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"golang_blog/config"
)

var DB *gorm.DB
var DBCasbin *gorm.DB

func init() {
	mysqlConfig := config.GetMysqlConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s", mysqlConfig.UserName, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Ip, mysqlConfig.DBName, mysqlConfig.Charset, mysqlConfig.ParseTime)
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
	if err != nil {
		panic(err)
	}
	dsnCasbin := fmt.Sprintf("%s:%s@tcp(%s:%s)/casbin?charset=%s&parseTime=%s", mysqlConfig.UserName, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Ip, mysqlConfig.Charset, mysqlConfig.ParseTime)
	dbConn1, err := gorm.Open(mysql.Open(dsnCasbin), &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
	if err != nil {
		panic(err)
	}
	DB = dbConn
	DBCasbin = dbConn1
}
