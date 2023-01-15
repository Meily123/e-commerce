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
// @Failure 400 {object}  ErrorResponse
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
// @Success 200 {object} LoginSuccessResponse
// @Failure      400  {object}  ErrorResponse
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
			"error": err,
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*2, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
		"token":   token,
	})

}

// SelfRequestUserHandler godoc
// @Summary get self request user
// @Description get logged in request user
// @Tags User
// @Produce  json
// @Param Cookie header string  false "token"
// @Success 200 {object} LoginSuccessResponse
// @Failure      400  {object}  ErrorResponse
// @Router /user/ [get]
func (userHandle *UserHandler) SelfRequestUserHandler(c *gin.Context) {
	user, err := c.Get("user")

	if err != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "user not authorized",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
		"user":    user,
	})
}

// SelfDeleteUserHandler godoc
// @Summary Delete self request user
// @Description Delete logged In self request user
// @Tags User
// @Produce  json
// @Param Cookie header string  false "token"
// @Success 200 {object} LoginSuccessResponse
// @Failure      400  {object}  ErrorResponse
// @Router /user/ [DELETE]
func (userHandle *UserHandler) SelfDeleteUserHandler(c *gin.Context) {

	user, err := c.Get("user")

	if err != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "user not authorized",
		})
		return
	}

	er := userHandle.userService.DeleteById(user.(model.User))

	if er != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "user not authorized",
		})
		return
	}

	fmt.Println("success?")

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
		"user":    user,
	})
}
