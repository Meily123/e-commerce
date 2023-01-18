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
	userVersionRoute.GET("/", middleware.RequireAuthentication, userHandler.SelfRequestUserHandler)
	userVersionRoute.DELETE("/", middleware.RequireAuthentication, userHandler.SelfDeleteUserHandler)
	userVersionRoute.PUT("/", middleware.RequireAuthentication, userHandler.SelfUpdateUserHandler)

	// user route, only admin Access
	adminUserVersionRoute := versionRoute.Group("/admin/user")
	adminUserVersionRoute.GET("/all", middleware.RequireAuthentication, middleware.AdminOnly, userHandler.GetAllUserHandler)
	adminUserVersionRoute.GET("/:id", middleware.RequireAuthentication, middleware.AdminOnly, userHandler.GetFindById)
	adminUserVersionRoute.PATCH("/:id", middleware.RequireAuthentication, middleware.AdminOnly, userHandler.UpdateUserToAdminHandler)
	adminUserVersionRoute.PUT("/:id", middleware.RequireAuthentication, middleware.AdminOnly, userHandler.UpdateUserHandler)

}
