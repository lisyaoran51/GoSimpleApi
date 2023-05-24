package controller

import (
	// "net/http"

	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lisyaoran51/GoSimpleApi/entity"
	"github.com/lisyaoran51/GoSimpleApi/service"
)

type UserController interface {
	Save(ctx *gin.Context) error
	Validate(ctx *gin.Context) error
}

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) UserController {
	return &userController{
		service: service,
	}
}

func (c *userController) Save(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}

	fmt.Printf("LOG: user: %#v\n", user)

	_, err = c.service.Save(user)

	return err
}

func (c *userController) Validate(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}

	return c.service.Validate(user)
}
