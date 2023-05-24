package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/lisyaoran51/GoSimpleApi/api"
	"github.com/lisyaoran51/GoSimpleApi/controller"
	"github.com/lisyaoran51/GoSimpleApi/repository"
	"github.com/lisyaoran51/GoSimpleApi/service"
)

var (
	userRepository repository.UserRepository = repository.NewMemoryUserRepository()
	userService    service.UserService       = service.NewUserService(userRepository)
	userController controller.UserController = controller.NewUserController(userService)
)

func main() {

	fmt.Println("start GoSimpleApi...")

	server := gin.New()

	server.Use(gin.Recovery(), gin.Logger())

	userAPI := api.NewUserAPI(userController)

	routes := server.Group("v1")
	{
		users := routes.Group("/users")
		{
			users.POST("create", userAPI.CreateUser)
			users.POST("validate", userAPI.Validate)
		}
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "5000"
	}

	server.Run(":" + port)
}
