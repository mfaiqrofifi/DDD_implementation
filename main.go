package main

import (
	repository "DDD/app/Repository"
	"DDD/app/handler"
	"DDD/app/services"
	"DDD/config"

	// "DDD/middleware"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db := config.InitDB()

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	e.POST("/users", userHandler.CreateUser)
	e.GET("/users/:id", userHandler.GetUser, handler.AuthMiddleWare)
	e.Logger.Fatal(e.Start(":8080"))
}
