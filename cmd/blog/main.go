package main

import (
	"Go_Day06/api/handler"
	"Go_Day06/config"
	"Go_Day06/models"
	"fmt"
	"log"
)

// @title Blog API
// @version 1.0
// @description API для управления блогом с JWT аутентификацией
// @host localhost:8080
// @BasePath /
func main() {
	cfg := config.GetConfig()
	fmt.Println(cfg)
	db, err := models.New(&cfg.DbConfig)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("ok")
	}
	// создаем модели
	userModel := models.NewUserModel(db)
	postModel := models.NewPostModel(db)
	h := handler.NewHandler(*userModel, *postModel)
	router := h.SetupRouters()
	if err := router.Run(":" + cfg.ServerConfig.Port); err != nil {
		log.Fatal("Ошибка запуска сервера: %v", err)
	}
}
