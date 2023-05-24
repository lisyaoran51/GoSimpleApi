package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lisyaoran51/GoSimpleApi/controller"
	"github.com/lisyaoran51/GoSimpleApi/dto"
)

type UserApi struct {
	userController controller.UserController
}

func NewUserAPI(userController controller.UserController) *UserApi {
	return &UserApi{
		userController: userController,
	}
}

// CreateUser godoc
// @Summary Create new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body entity.User true "Create user"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /users [post]
func (api *UserApi) CreateUser(ctx *gin.Context) {

	err := api.userController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Reason: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Success: true,
		})
	}
}

func (api *UserApi) Validate(ctx *gin.Context) {

	err := api.userController.Validate(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Reason: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Success: true,
		})
	}
}
