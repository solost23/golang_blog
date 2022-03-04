package mysql

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"jwt-go/config"
)

var DB *gorm.DB

func init() {
	mysqlConfig := config.GetMysqlConfig()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysqlConfig.UserName, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Ip, mysqlConfig.DBName)
	DB, _ = gorm.Open("mysql", dsn)

	DB.Debug()
	DB.LogMode(true)

	if err := DB.DB().Ping(); err != nil {
		log.Fatalln(err)
	}
}
