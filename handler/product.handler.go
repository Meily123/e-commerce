package handler

import (
	"WebAPI/model"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

// ProductHandler godoc
// @Summary get product
// @Description get detail product
// @Tags Product
// @Param id path int true "uuid"
// @Accept  json
// @Produce  json
// @Success 200  {object} model.Product
// @Router /product/{id} [get]
func ProductHandler(c *gin.Context) {
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
// @Param Body body ProductInput true "Product"
// @Accept  json
// @Produce  json
// @Success 200  {object} model.Product
// @Router /product [post]
func ProductInputHandler(c *gin.Context) {
	var productInput model.ProductInput

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
}
