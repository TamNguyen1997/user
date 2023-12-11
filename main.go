package main

import (
	"github.com/gin-gonic/gin"
	"main.go/controller"
	"main.go/model"
	"main.go/repository"
	"main.go/service"
)

func main() {
	db := model.ConnectPostgreSql()
	userRepository := repository.NewUserRepo(db)
	userService := service.NewUserService(userRepository)
	c := controller.NewUserController(userService)

	server := gin.Default()
	c.AddRoutes(server)

	server.Run()
}
