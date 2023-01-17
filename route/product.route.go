package route

import (
	"WebAPI/handler"
	"WebAPI/middleware"
	"WebAPI/repository"
	"WebAPI/service"
	"github.com/gin-gonic/gin"
)

func ProductRoute(versionRoute *gin.RouterGroup) {
	productRepository := repository.NewProductRepository()
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewProductHandler(productService)

	productVersionRoute := versionRoute.Group("/product")

	productVersionRoute.DELETE("/:id", middleware.RequireAuthentication, middleware.AdminOnly, productHandler.DeleteProductHandler)
	productVersionRoute.PUT("/:id", middleware.RequireAuthentication, middleware.AdminOnly, productHandler.UpdateProductHandler)
	productVersionRoute.GET("/:id", middleware.RequireAuthentication, productHandler.GetFindById)
	productVersionRoute.POST("", middleware.RequireAuthentication, middleware.AdminOnly, productHandler.CreateProductHandler)
	productVersionRoute.GET("", middleware.RequireAuthentication, productHandler.GetAllProductHandler)
}
