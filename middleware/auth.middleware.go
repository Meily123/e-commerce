package middleware

import (
	"WebAPI/model"
	"WebAPI/repository"
	"WebAPI/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"time"
)

func RequireAuthentication(c *gin.Context) {

	// Get the cookie
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode/Validate
	// Parse takes the token string and a function for looking up the key. The latter is especially
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check expire

		if float64(time.Now().UnixNano()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Find the user by sub token which contain User Id

		userRepository := repository.NewUserRepository()
		userService := service.NewUserService(userRepository)
		user, err := userService.FindById(claims["sub"].(string))

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		emptyUser := model.User{}

		if user == emptyUser {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach Req
		c.Set("user", user)
		c.Set("is_admin", user.IsAdmin)

		//Continue
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
