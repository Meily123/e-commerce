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
		return
	}

	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// Decode/Validate
	// Parse takes the token string and a function for looking up the key. The latter is especially
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check expire

		if float64(time.Now().UnixNano()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Find the user by sub token which contain User Id

		userRepository := repository.NewUserRepository()
		userService := service.NewUserService(userRepository)
		user, err := userService.FindById(claims["sub"].(string))

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if user.Id.String() == "00000000-0000-0000-0000-000000000000" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
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

func AdminOnly(c *gin.Context) {
	user, err := c.Get("user")

	if err != true {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":  401,
			"error": "user not authorized",
		})
		c.Abort()
		return
	}

	if user.(model.User).IsAdmin == false {
		c.JSON(http.StatusForbidden, gin.H{
			"code":  403,
			"error": "you don't have access",
		})
		c.Abort()
		return
	}
	return
}
