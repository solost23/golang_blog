package config

import (
	"log"

	"github.com/spf13/viper"
)

type Mysql struct {
	UserName string
	Password string
	Host     string
	Ip       string
	DBName   string
}

func NewMysql(userName, password, host, ip, dbName string) *Mysql {
	return &Mysql{
		UserName: userName,
		Password: password,
		Host:     host,
		Ip:       ip,
		DBName:   dbName,
	}
}

func GetMysqlConfig() *Mysql {
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("config read error")
	}

	username := viper.GetStringMapString("mysql")["username"]
	password := viper.GetStringMapString("mysql")["password"]
	host := viper.GetStringMapString("mysql")["host"]
	ip := viper.GetStringMapString("mysql")["ip"]
	dbname := viper.GetStringMapString("mysql")["dbname"]
	//fmt.Println(username, password, host, ip, dbname)
	mysql := NewMysql(username, password, host, ip, dbname)
	return mysql
}
