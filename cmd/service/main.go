package main

import (
	"THackaton/internal/application/service"
	"THackaton/internal/infrastructure/database"
	"THackaton/internal/infrastructure/repository"
	"THackaton/internal/infrastructure/transport/http/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewSQLiteDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Сборка зависимостей
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// HTTP-сервер
	r := gin.Default()
	r.POST("/register", userHandler.Register)

	r.Run(":8080")
}
