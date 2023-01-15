package route

import (
	"github.com/gin-gonic/gin"
)

func UserRoute(versionRoute *gin.RouterGroup) {
	userVersionRoute := versionRoute.Group("/user")

	userVersionRoute.GET("/")
	userVersionRoute.DELETE("/")
	userVersionRoute.GET("/all")
	userVersionRoute.GET("/:id")
	userVersionRoute.PATCH("/admin/:id")
	userVersionRoute.PUT("/edit/:id")
	userVersionRoute.PUT("/edit")
	userVersionRoute.GET("/edit")
}
