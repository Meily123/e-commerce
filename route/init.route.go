package route

import "github.com/gin-gonic/gin"

func InitRoute(router *gin.Engine, version string) {

	// group version router
	versionRouter := router.Group(version)

	//routes
	ProductRoute(versionRouter)
	UserRoute(versionRouter)
	CartRoute(versionRouter)
}
