package route

import (
	"WebAPI/handler"
	"github.com/gin-gonic/gin"
)

func ProductRoute(versionRoute *gin.RouterGroup) {
	versionRoute.GET("/product/:id", handler.ProductHandler)
	versionRoute.POST("/product", handler.ProductInputHandler)
	versionRoute.GET("/product")
}
