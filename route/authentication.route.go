package route

import (
	"WebAPI/handler"
	"WebAPI/repository"
	"WebAPI/service"
	"github.com/gin-gonic/gin"
)

func AuthenticationRoute(versionRoute *gin.RouterGroup) {
	// user handler init
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	versionRoute.POST("/register", userHandler.RegistrationHandler)
	versionRoute.POST("/login")
}
