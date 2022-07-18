package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang_blog/models"
	"golang_blog/router"
)

// @title Blog Swagger
// @version 1.0
// @description This is a blog
// @host localhost:8080
// @schemes http https
// @BasePath /
func main() {
	// 创建数据库表
	InitAutoMigrate()
	s := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      router.Register(),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	fmt.Printf("博客服务开启, 服务地址:%s", "0.0.0.0:8080 \n")
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err.Error())
	}
}

func InitAutoMigrate() {
	var db = models.DB
	err := db.AutoMigrate(&models.User{}, &models.Category{}, &models.Article{}, &models.Comment{}, &models.CasbinModel{})
	if err != nil {
		panic(err)
	}
}
