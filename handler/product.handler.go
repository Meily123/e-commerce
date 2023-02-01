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
// @Success 200 {object} SuccessResponse{data=model.ProductResponse}
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

	// handle error saving product
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
		"product": product,
	})
}

// DeleteProductHandler godoc
// @Summary Delete product
// @Description Delete product by id
// @Tags Product
// @Produce  json
// @Param id path string true "uuid"
// @Param Cookie header string  false "token"
// @Success 200 {object} BaseSuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /product/{id} [DELETE]
func (productHandle *ProductHandler) DeleteProductHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	err := productHandle.productService.DeleteById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
	})
}

// GetAllProductHandler godoc
// @Summary Get all products
// @Description Get all products data
// @Tags Product
// @Produce  json
// @Param Cookie header string  false "token"
// @Success 200 {object} SuccessResponse{data=[]model.ProductResponse}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /product [GET]
func (productHandle *ProductHandler) GetAllProductHandler(c *gin.Context) {

	products, err := productHandle.productService.FindAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
		"product": products,
	})
}

// GetFindById godoc
// @Summary Get product by id
// @Description Get product base on id parameters given
// @Tags Product
// @Produce  json
// @Param id path string true "uuid"
// @Param Cookie header string  false "token"
// @Success 200 {object} SuccessResponse{data=model.ProductResponse}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /product/{id} [GET]
func (productHandle *ProductHandler) GetFindById(c *gin.Context) {
	id := c.Params.ByName("id")

	product, err := productHandle.productService.GetFindById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
		"product": product,
	})
}

// UpdateProductHandler godoc
// @Summary Update product
// @Description Update product data (Admin Only)
// @Tags Product
// @Produce  json
// @Param id path string true "uuid"
// @Param Body body ProductEditRequest true "Product"
// @Param Cookie header string  false "token"
// @Success 200 {object} SuccessResponse{data=model.ProductResponse}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /product/{id} [PUT]
func (productHandle *ProductHandler) UpdateProductHandler(c *gin.Context) {
	// body
	var editRequest model.ProductEditRequest
	err := c.BindJSON(&editRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": err.Error(),
		})
		return
	}

	// param id
	id := c.Params.ByName("id")

	product, err := productHandle.productService.UpdateProduct(id, editRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
		"product": product,
	})
}
