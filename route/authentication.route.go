package route

import (
	"WebAPI/handler"
	"github.com/gin-gonic/gin"
)

func AuthenticationRoute(versionRoute *gin.RouterGroup) {
	versionRoute.POST("/register", handler.RegistrationHandler)
	versionRoute.POST("/login")
}
