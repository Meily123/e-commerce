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

type CartHandler struct {
	cartService service.CartService
}

func NewCartHandler(cartServ service.CartService) *CartHandler {
	return &CartHandler{cartServ}
}

// CreateCartHandler godoc
// @Summary Create Cart
// @Description Add product to cart
// @Tags Cart
// @Param Body body CartRequest true "cart"
// @Param Cookie header string  false "token"
// @Accept  json
// @Produce  json
// @success 200 {object} model.CartProduct
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /cart [post]
func (cartHandle *CartHandler) CreateCartHandler(c *gin.Context) {
	user, isExists := c.Get("user")

	if isExists != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "user not authorized",
		})
		return
	}

	var cartRequest model.CartProductRequest

	err := c.BindJSON(&cartRequest)

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

	cart, err := cartHandle.cartService.CreateCart(cartRequest, user.(model.User))

	// handle error saving cart
	if err != nil {
		fmt.Println("here")
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
		"cart":    cart,
	})
}

// DeleteCartHandler godoc
// @Summary Delete cart
// @Description Delete cart by id
// @Tags Cart
// @Produce  json
// @Param id path string true "uuid"
// @Param Cookie header string  false "token"
// @success 200 {object} CartResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /cart/{id} [DELETE]
func (cartHandle *CartHandler) DeleteCartHandler(c *gin.Context) {
	user, isExists := c.Get("user")

	if isExists != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "user not authorized",
		})
		return
	}

	id := c.Params.ByName("id")

	er := cartHandle.cartService.DeleteById(id, user.(model.User))

	if er != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  403,
			"error": "you don't have access",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
	})
}

// GetAllCartHandler godoc
// @Summary Get all carts
// @Description Get all carts data
// @Tags Cart
// @Produce  json
// @Param Cookie header string  false "token"
// @success 200 {array} CartResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /cart [GET]
func (cartHandle *CartHandler) GetAllCartHandler(c *gin.Context) {
	user, isExists := c.Get("user")

	if isExists != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "user not authorized",
		})
		return
	}

	carts, er := cartHandle.cartService.FindAll(user.(model.User))

	if er != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  403,
			"error": "you don't have access",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
		"cart":    carts,
	})
}

// UpdateCartHandler godoc
// @Summary Update cart
// @Description Update cart data (Admin Only)
// @Tags Cart
// @Produce  json
// @Param id path string true "uuid"
// @Param Body body CartEditRequest true "Cart"
// @Param Cookie header string  false "token"
// @success 200 {object} CartResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /cart/{id} [PUT]
func (cartHandle *CartHandler) UpdateCartHandler(c *gin.Context) {
	user, isExists := c.Get("user")

	if isExists != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "user not authorized",
		})
		return
	}

	// body
	var editRequest model.CartProductEditRequest
	err := c.BindJSON(&editRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  500,
			"error": err,
		})
		return
	}

	// param id
	id := c.Params.ByName("id")

	cart, err := cartHandle.cartService.UpdateCart(id, editRequest, user.(model.User))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": "error bad request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
		"cart":    cart,
	})
}
