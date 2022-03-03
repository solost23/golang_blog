package main

import (
	"log"
	"net/http"
	"time"

	"jwt-go/model"
	"jwt-go/router"
)

func main() {
	var DB = model.DB
	DB.AutoMigrate(&model.User{})
	s := &http.Server{
		Addr:         "localhost:8080",
		Handler:      router.Register(),
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}
