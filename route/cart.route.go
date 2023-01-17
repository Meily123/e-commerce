package route

import (
	"WebAPI/handler"
	"WebAPI/middleware"
	"WebAPI/repository"
	"WebAPI/service"
	"github.com/gin-gonic/gin"
)

func CartRoute(versionRoute *gin.RouterGroup) {
	cartRepository := repository.NewCartRepository()
	cartService := service.NewCartService(cartRepository)
	cartHandler := handler.NewCartHandler(cartService)

	cartVersionRoute := versionRoute.Group("/cart")

	cartVersionRoute.DELETE("/:id", middleware.RequireAuthentication, middleware.AdminOnly, cartHandler.DeleteCartHandler)
	cartVersionRoute.PUT("/:id", middleware.RequireAuthentication, middleware.AdminOnly, cartHandler.UpdateCartHandler)
	cartVersionRoute.POST("", middleware.RequireAuthentication, middleware.AdminOnly, cartHandler.CreateCartHandler)
	cartVersionRoute.GET("", middleware.RequireAuthentication, cartHandler.GetAllCartHandler)
}
