package handler

import (
	"WebAPI/model"
	"WebAPI/service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type ProductHandler struct {
	productService service.ProductService
}

func NewProductHandler(productServ service.ProductService) *ProductHandler {
	return &ProductHandler{productServ}
}

// CreateProductHandler godoc
// @Summary Create Product
// @Description Create New Product
// @Tags Product
// @Param Body body ProductRequest true "product"
// @Param Cookie header string  false "token"
// @Accept  json
// @Produce  json
// @success 200 {object} model.Product
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /product [post]
func (productHandle *ProductHandler) CreateProductHandler(c *gin.Context) {
	var productRequest model.ProductRequest

	err := c.BindJSON(&productRequest)

	// handle error binding and validation
	if err != nil {
		var ve validator.ValidationErrors
		var errorMassages []string

		if errors.As(err, &ve) {
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("field: %s, condition: %s", e.Field(), e.ActualTag())
				errorMassages = append(errorMassages, errorMessage)
			}
		} else {
			errorMassages = append(errorMassages, err.Error())
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": errorMassages,
		})
		return
	}

	product, err := productHandle.productService.CreateProduct(productRequest)

	// handle error saving user
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
		"product": product,
	})
}

/*
// ProductHandler1 godoc
// @Summary get product
// @Description get detail product
// @Tags Product
// @Param id path int true "uuid"
// @Accept  json
// @Produce  json
// @Success 200  {object} model.Product
// @Router /product/{id} [get]
func ProductHandler1(c *gin.Context) {
	id := c.Params.ByName("id")
	//id := c.Query("id")
	//prod := model.Product{id}
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"judul": "title",
	})
}

// ProductInputHandler godoc
// @Summary post product
// @Description post create product
// @Tags Product
// @Param Body body ProductRequest true "Product"
// @Accept  json
// @Produce  json
// @Success 200  {object} model.Product
// @Router /product [post]
func ProductInputHandler(c *gin.Context) {
	var productInput model.ProductRequest

	err := c.BindJSON(&productInput)
	if err != nil {
		var ve validator.ValidationErrors
		var errorMassages []string
		if errors.As(err, &ve) {
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
				errorMassages = append(errorMassages, errorMessage)
			}
		} else {
			errorMassages = append(errorMassages, err.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"massage": errorMassages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":        productInput.Name,
		"base_price":   productInput.BasePrice,
		"sell_price":   productInput.SellPrice,
		"stock":        productInput.Stock,
		"descriptions": productInput.Description,
	})
}*/
