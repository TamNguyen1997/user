package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"main.go/model"
	"main.go/service"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

func (c *UserController) findAll(ctx *gin.Context) {
	result, err := c.userService.FindAll()
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, result)
}

func (c *UserController) find(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	u, err := c.userService.Find(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, u)
}

func (c *UserController) delete(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	u, err := c.userService.Find(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, u)
}

func (c *UserController) save(ctx *gin.Context) {
	var result *model.User

	if err := ctx.ShouldBindWith(&result, binding.JSON); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	result, err := c.userService.Save(result)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	ctx.IndentedJSON(http.StatusOK, result)
}

func (c *UserController) AddRoutes(ctx *gin.Engine) {
	route := ctx.Group("/users/")
	{
		route.GET("", c.findAll)
		route.POST("", c.save)
		route.GET(":id", c.find)
		route.DELETE(":id", c.delete)
	}
}
