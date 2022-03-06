package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"jwt-go/model"
	"jwt-go/router"
)

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
