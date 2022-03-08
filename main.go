package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"golang_blog/model"
	"golang_blog/router"
)

// @title Blog Swagger
// @version 1.0
// @description This is a blog
// @host localhost:8080
// @schemes http https
// @BasePath /
func main() {
	var DB = model.DB
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.Content{})
	DB.AutoMigrate(&model.Article{})
	s := &http.Server{
		Addr:         "localhost:8080",
		Handler:      router.Register(),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	fmt.Printf("博客服务开启, 服务地址:%s", "localhost:8080")
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
