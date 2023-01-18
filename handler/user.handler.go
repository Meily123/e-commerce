package handler

import (
	"WebAPI/model"
	"WebAPI/service"
	"WebAPI/shared"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
// @Description Create User
// @Tags Authentication
// @Param Body body UserRequest true "User"
// @Accept  json
// @Produce  json
// @Success 200 {object} SuccessResponse{data=model.UserResponse}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
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
			"error": "something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"massage": "success",
		"user":    shared.UserRenderToResponse(user),
	})
}

// LoginHandler godoc
// @Summary Login
// @Description Login User
// @Tags Authentication
// @Param Body body LoginRequest true "User"
// @Accept  json
// @Produce  json
// @Success 200 {object} SuccessResponse{data=string}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /login [post]
func (userHandle *UserHandler) LoginHandler(c *gin.Context) {

	// user login input
	var loginRequest model.LoginRequest
	err := c.BindJSON(&loginRequest)

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

	//find user by username
	token, err := userHandle.userService.LoginUser(loginRequest)

	// handle find saving user
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":  401,
			"error": "invalid credentials wrong username or password",
		})

		return
	}

	// set token to cookie
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
// @Success 200 {object} SuccessResponse{data=model.UserResponse}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
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
		"user":    shared.UserRenderToResponse(user.(model.User)),
	})
}

// SelfDeleteUserHandler godoc
// @Summary Delete self request user
// @Description Delete logged In self request user
// @Tags User
// @Produce  json
// @Param Cookie header string  false "token"
// @Success 200 {object} SuccessResponse{data=model.UserResponse}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/ [DELETE]
func (userHandle *UserHandler) SelfDeleteUserHandler(c *gin.Context) {

	user, exist := c.Get("user")

	if exist != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "user not authorized",
		})
		return
	}

	err := userHandle.userService.DeleteById(user.(model.User))

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
		"user":    shared.UserRenderToResponse(user.(model.User)),
	})
}

// GetAllUserHandler godoc
// @Summary Get all users
// @Description Get all users data (Admin Only)
// @Tags User
// @Produce  json
// @Param Cookie header string  false "token"
// @Success 200 {object} SuccessResponse{data=[]model.UserResponse}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/all [GET]
func (userHandle *UserHandler) GetAllUserHandler(c *gin.Context) {

	users, err := userHandle.userService.FindAll()

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
		"user":    shared.ListUserRenderToResponse(users),
	})
}

// GetFindById godoc
// @Summary Get user by id
// @Description Get user base on id parameters given (Admin Only)
// @Tags User
// @Produce  json
// @Param id path string true "uuid"
// @Param Cookie header string  false "token"
// @Success 200 {object} SuccessResponse{data=model.UserResponse}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/admin/{id} [GET]
func (userHandle *UserHandler) GetFindById(c *gin.Context) {
	id := c.Params.ByName("id")

	user, err := userHandle.userService.GetFindById(id)

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
		"user":    shared.UserRenderToResponse(user),
	})
}

// UpdateUserToAdminHandler godoc
// @Summary Update user to admin
// @Description Update user into admin by id (Admin Only)
// @Tags User
// @Produce  json
// @Param id path string true "uuid"
// @Param Cookie header string  false "token"
// @Success 200 {object} SuccessResponse{data=model.UserResponse}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/admin/{id} [PATCH]
func (userHandle *UserHandler) UpdateUserToAdminHandler(c *gin.Context) {
	id := c.Params.ByName("id")

	user, err := userHandle.userService.UpdateAdminById(id)

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
		"user":    shared.UserRenderToResponse(user),
	})
}

// UpdateUserHandler godoc
// @Summary Update user
// @Description Update user data by id user (Admin Only)
// @Tags User
// @Produce  json
// @Param id path string true "uuid"
// @Param Body body UserEditRequest true "User"
// @Param Cookie header string  false "token"
// @Success 200 {object} SuccessResponse{data=model.UserResponse}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/admin/{id} [PUT]
func (userHandle *UserHandler) UpdateUserHandler(c *gin.Context) {
	// body
	var editRequest model.UserEditRequest

	err := c.BindJSON(&editRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": "bad request",
		})
		return
	}

	// param id
	id := c.Params.ByName("id")

	user, err := userHandle.userService.UpdateUser(id, editRequest)

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
		"user":    shared.UserRenderToResponse(user),
	})
}

// SelfUpdateUserHandler godoc
// @Summary Update self user
// @Description Update self request user data
// @Tags User
// @Produce  json
// @Param Body body UserEditRequest true "User"
// @Param Cookie header string  false "token"
// @Success 200 {object} SuccessResponse{data=model.UserResponse}
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /user/ [PUT]
func (userHandle *UserHandler) SelfUpdateUserHandler(c *gin.Context) {
	// body
	var editRequest model.UserEditRequest
	err := c.BindJSON(&editRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":  400,
			"error": "bad request",
		})
		return
	}

	// get user
	selfRequestUser, isExists := c.Get("user")

	if isExists != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "user not authorized",
		})
		return
	}

	user, err := userHandle.userService.SelfUpdateUser(selfRequestUser.(model.User), editRequest)

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
		"user":    shared.UserRenderToResponse(user),
	})
}
