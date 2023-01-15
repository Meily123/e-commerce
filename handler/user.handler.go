package handler

import (
	"WebAPI/model"
	"WebAPI/service"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"net/http"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userServ service.UserService) *UserHandler {
	return &UserHandler{userServ}
}

// RegistrationHandler godoc
// @Summary Register
// @Description Register User
// @Tags User
// @Param Body body UserRequest true "User"
// @Accept  json
// @Produce  json
// @Success 200  {object} model.User
// @Router /register [post]
func (userHandle *UserHandler) RegistrationHandler(c *gin.Context) {
	var userRequest model.UserRequest

	err := c.BindJSON(&userRequest)

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

	user, err := userHandle.userService.CreateUser(userRequest)

	// handle error saving user
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
		"user":    user,
	})
}

// LoginHandler godoc
// @Summary Register
// @Description Register User
// @Tags User
// @Param Body body LoginRequest true "User"
// @Accept  json
// @Produce  json
// @Success 200  {object} model.User
// @Router /login [post]
func (userHandle *UserHandler) LoginHandler(c *gin.Context) {

	// user login input
	var loginRequest model.LoginRequest
	err := c.BindJSON(&loginRequest)

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

	//find user by username
	token, err := userHandle.userService.LoginUser(loginRequest)
	fmt.Println(token)
	// handle find saving user
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":  401,
				"error": "record not found",
			})
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": "this2",
		})

		return
	}

	//return if not empty
	if token == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  500,
			"error": "tis1",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
		"token":   token,
	})

}
