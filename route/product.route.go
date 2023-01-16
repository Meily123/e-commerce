package route

import (
	"github.com/gin-gonic/gin"
)

func ProductRoute(versionRoute *gin.RouterGroup) {
	productVersionRoute := versionRoute.Group("/product")

	productVersionRoute.GET("/:id")
	productVersionRoute.DELETE("/:id")
	productVersionRoute.PUT("/:id")
	productVersionRoute.POST("/")
	productVersionRoute.GET("/")
}
