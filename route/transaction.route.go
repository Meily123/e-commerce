package route

import (
	"WebAPI/handler"
	"WebAPI/middleware"
	"WebAPI/repository"
	"WebAPI/service"
	"github.com/gin-gonic/gin"
)

func TransactionRoute(versionRoute *gin.RouterGroup) {
	transactionRepository := repository.NewTransactionRepository()
	transactionService := service.NewTransactionService(transactionRepository)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	transactionVersionRoute := versionRoute.Group("/transaction")
	transactionVersionRoute.GET("", middleware.RequireAuthentication, transactionHandler.DetailTransactionHandler)
	transactionVersionRoute.POST("", middleware.RequireAuthentication, transactionHandler.CreateTransactionHandler)
	transactionVersionRoute.GET("/:id", middleware.RequireAuthentication, transactionHandler.SelfByIdTransactionHandler)
	transactionVersionRoute.GET("/all", middleware.RequireAuthentication, transactionHandler.SelfAllTransactionHandler)

	adminTransactionVersionRoute := versionRoute.Group("/admin/transaction")
	adminTransactionVersionRoute.GET("/all", middleware.RequireAuthentication, middleware.AdminOnly, transactionHandler.AllTransactionHandler)
	adminTransactionVersionRoute.PATCH("/:id", middleware.RequireAuthentication, middleware.AdminOnly, transactionHandler.VerifyPaymentTransactionHandler)
	adminTransactionVersionRoute.GET("/:id", middleware.RequireAuthentication, middleware.AdminOnly, transactionHandler.FindByIdTransactionHandler)
	adminTransactionVersionRoute.GET("/summary", middleware.RequireAuthentication, middleware.AdminOnly, transactionHandler.SummaryTransactionHandler)

}
