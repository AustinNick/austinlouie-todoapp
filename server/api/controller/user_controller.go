package controller

import (
	"server/api/service"

	"github.com/gin-gonic/gin"
)

type UserController interface {
	LoginHandler(ctx *gin.Context)
	RegisterHandler(ctx *gin.Context)
	LogoutHandler(ctx *gin.Context)
	RefreshHandler(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{
		userService: userService,
	}
}

func (c *userController) LoginHandler(ctx *gin.Context) {
	// TODO
}

func (c *userController) RegisterHandler(ctx *gin.Context) {
	// TODO
}

func (c *userController) LogoutHandler(ctx *gin.Context) {
	// TODO
}

func (c *userController) RefreshHandler(ctx *gin.Context) {
	// TODO
}
