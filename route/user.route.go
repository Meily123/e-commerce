package route

import (
	"WebAPI/handler"
	"WebAPI/middleware"
	"WebAPI/repository"
	"WebAPI/service"
	"github.com/gin-gonic/gin"
)

func UserRoute(versionRoute *gin.RouterGroup) {
	// user handler init
	userRepository := repository.NewUserRepository()
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// user auth
	versionRoute.POST("/register", userHandler.RegistrationHandler)
	versionRoute.POST("/login", userHandler.LoginHandler)

	// user group route
	userVersionRoute := versionRoute.Group("/user")

	// user routes
	userVersionRoute.GET("/", middleware.RequireAuthentication, userHandler.RequestUserHandler)
	userVersionRoute.DELETE("/")
	userVersionRoute.GET("/all")
	userVersionRoute.GET("/:id")
	userVersionRoute.PATCH("/admin/:id")
	userVersionRoute.PUT("/edit/:id")
	userVersionRoute.PUT("/edit")
	userVersionRoute.GET("/edit")
}
