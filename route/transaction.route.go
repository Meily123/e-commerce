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

}
