package handler

import (
	"WebAPI/model"
	"WebAPI/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionServ service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionServ}
}

// CreateTransactionHandler godoc
// @Summary Create Transaction
// @Description Create New Transaction
// @Tags Transaction
// @Param Cookie header string  false "token"
// @Accept  json
// @Produce  json
// @success 200 {object} model.TransactionProduct
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transaction [get]
func (transactionHandle *TransactionHandler) CreateTransactionHandler(c *gin.Context) {
	user, isExists := c.Get("user")

	if isExists != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "user not authorized",
		})
		return
	}

	transaction, err := transactionHandle.transactionService.CreateTransaction(user.(model.User))

	// handle error saving transaction
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":        200,
		"massage":     "success",
		"transaction": transaction,
	})
}

// DetailTransactionHandler godoc
// @Summary Create Transaction
// @Description Create New Transaction
// @Tags Transaction
// @Param Cookie header string  false "token"
// @Accept  json
// @Produce  json
// @success 200 {object} model.TransactionProduct
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transaction [post]
func (transactionHandle *TransactionHandler) DetailTransactionHandler(c *gin.Context) {
	user, isExists := c.Get("user")

	if isExists != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "user not authorized",
		})
		return
	}

	transaction, err := transactionHandle.transactionService.DetailTransaction(user.(model.User))

	// handle error saving transaction
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":        200,
		"massage":     "success",
		"transaction": transaction,
	})
}
