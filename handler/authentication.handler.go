package handler

import (
	"WebAPI/config"
	"WebAPI/model"
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
func RegistrationHandler(c *gin.Context) {
	var userInput model.UserRequest

	err := c.BindJSON(&userInput)
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

	user := model.User{
		Name:     userInput.Name,
		Username: userInput.UserName,
		Address:  userInput.Address,
		Password: userInput.Password,
		Email:    userInput.Email,
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   500,
			"massage": "database error registering user",
		})
	}

	result := config.ConnectToDatabase().Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   500,
			"massage": "database error registering user",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"massage": "success",
	})
}
