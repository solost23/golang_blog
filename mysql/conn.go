package mysql

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	dsn := "root:123@tcp(localhost:3306)/jwt-go"
	DB, _ = gorm.Open("mysql", dsn)

	DB.Debug()
	DB.LogMode(true)

	if err := DB.DB().Ping(); err != nil {
		log.Fatalln(err)
	}
}
