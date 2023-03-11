package injector

import (
	"server/api/controller"

	"server/api/repository"
	"server/api/service"

	"gorm.io/gorm"
)

func InitAuth(db *gorm.DB) controller.UserController {
	userRepository := repository.NewUserRepository(db)
	authService := service.NewUserService(userRepository)
	jwtService := service.NewJWTService()
	userController := controller.NewUserController(authService, jwtService)

	return userController
}
