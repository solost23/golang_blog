package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"golang_blog/config"
)

var DB *gorm.DB

func init() {
	mysqlConfig := config.GetMysqlConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlConfig.UserName, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Ip, mysqlConfig.DBName)
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Error)})
	if err != nil {
		panic(err)
	}
	DB = dbConn
}
