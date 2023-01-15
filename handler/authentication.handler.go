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

// RegistrationHandler godoc
// @Summary Register
// @Description Register User
// @Tags User
// @Param Body body UserRequest true "User"
// @Accept  json
// @Produce  json
// @Success 200  {object} UserResponse
// @Router /register [post]

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userServ service.UserService) *userHandler {
	return &userHandler{userServ}
}

func (userHandle userHandler) RegistrationHandler(c *gin.Context) {
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
			"error":   400,
			"massage": errorMassages,
		})
		return
	}

	user, err := userHandle.userService.CreateUser(userRequest)

	// handle error saving user
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   500,
			"massage": "database error registering user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"massage": "success",
		"user":    user,
	})
}
