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
// @success 200 {object} SuccessResponse{data=model.Transaction}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transaction [post]
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
		"code":    200,
		"massage": "success",
		"data":    transaction,
	})
}

// DetailTransactionHandler godoc
// @Summary Detail Transaction
// @Description Get New transaction detail
// @Tags Transaction
// @Param Cookie header string  false "token"
// @Accept  json
// @Produce  json
// @success 200 {object} SuccessResponse{data=model.Transaction}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transaction [get]
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
		"code":    200,
		"massage": "success",
		"data":    transaction,
	})
}

// VerifyPaymentTransactionHandler godoc
// @Summary Verify Transaction
// @Description Verify transaction detail
// @Tags Transaction
// @Param id path string true "uuid"
// @Param Cookie header string  false "token"
// @Accept  json
// @Produce  json
// @success 200 {object} BaseSuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/transaction/{id} [patch]
func (transactionHandle *TransactionHandler) VerifyPaymentTransactionHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	err := transactionHandle.transactionService.VerifyPaymentTransaction(id)

	// handle error saving transaction
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
	})
}

// AllTransactionHandler godoc
// @Summary Get  All Transaction
// @Description Get All transaction detail
// @Tags Transaction
// @Param Cookie header string  false "token"
// @Accept  json
// @Produce  json
// @success 200 {object} SuccessResponse{data=[]model.Transaction}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/transaction/all [get]
func (transactionHandle *TransactionHandler) AllTransactionHandler(c *gin.Context) {

	transactions, err := transactionHandle.transactionService.AllTransaction()

	// handle error saving transaction
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
		"data":    transactions,
	})
}

// FindByIdTransactionHandler godoc
// @Summary Get Transaction by id
// @Description Get transaction detail by id
// @Tags Transaction
// @Param Cookie header string  false "token"
// @Accept  json
// @Produce  json
// @success 200 {object} SuccessResponse{data=model.Transaction}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/transaction/{id} [get]
func (transactionHandle *TransactionHandler) FindByIdTransactionHandler(c *gin.Context) {
	user, isExists := c.Get("user")

	if isExists != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "user not authorized",
		})
		return
	}

	id := c.Params.ByName("id")

	transaction, err := transactionHandle.transactionService.FindByIdTransaction(id, user.(model.User))

	// handle error saving transaction
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
		"data":    transaction,
	})
}

// SelfAllTransactionHandler godoc
// @Summary Get All Own Transaction
// @Description Get All Self Own Detail Transaction
// @Tags Transaction
// @Param Cookie header string  false "token"
// @Accept  json
// @Produce  json
// @success 200 {object} SuccessResponse{data=[]model.Transaction}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transaction/all [get]
func (transactionHandle *TransactionHandler) SelfAllTransactionHandler(c *gin.Context) {
	user, isExists := c.Get("user")

	if isExists != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "user not authorized",
		})
		return
	}

	transactions, err := transactionHandle.transactionService.SelfAllTransaction(user.(model.User))

	// handle error saving transaction
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
		"data":    transactions,
	})
}

// SelfByIdTransactionHandler godoc
// @Summary Get By Id Own Transaction
// @Description Get By Id Self Own Detail Transaction
// @Tags Transaction
// @Param Cookie header string  false "token"
// @Accept  json
// @Produce  json
// @success 200 {object} SuccessResponse{data=model.Transaction}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /transaction/{id} [get]
func (transactionHandle *TransactionHandler) SelfByIdTransactionHandler(c *gin.Context) {
	user, isExists := c.Get("user")

	if isExists != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "user not authorized",
		})
		return
	}

	id := c.Params.ByName("id")

	transaction, err := transactionHandle.transactionService.FindByIdTransaction(id, user.(model.User))

	// handle error saving transaction
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
		"data":    transaction,
	})
}

// SummaryTransactionHandler godoc
// @Summary Get Summary Sold Product
// @Description Get Summary of All Sold Product
// @Tags Transaction
// @Param Cookie header string  false "token"
// @Accept  json
// @Produce  json
// @success 200 {object} SuccessResponse{data=model.TransactionSummaryResponse}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /admin/transaction/summary [get]
func (transactionHandle *TransactionHandler) SummaryTransactionHandler(c *gin.Context) {
	transactionSummary, err := transactionHandle.transactionService.SummaryTransaction()

	// handle error saving transaction
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
		"data":    transactionSummary,
	})
}
