package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ilyasbabu/go-ecommerce/controllers"
	"github.com/ilyasbabu/go-ecommerce/models"
	"gorm.io/gorm"
)

func RequireAuth(c *gin.Context) {
	r := controllers.Response()
	tokenString := c.PostForm("Token")
	if tokenString == "" {
		fmt.Println("Failed to retrieve token")
		r.Message = "Authorization Failed"
		c.JSON(http.StatusNotAcceptable, r)
		c.Abort()
	}

	// tokenString, err := c.Cookie("Authorization")
	// if err != nil {
	// 	r.Message = "Authorization Failed"
	// 	c.JSON(http.StatusForbidden, r)
	// 	c.Abort()
	// }

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["sub"], claims["exp"])

		var user models.User
		err := controllers.Db.First(&user, claims["sub"]).Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Failed to retrieve user")
			r.Message = "Authorization Failed"
			c.JSON(http.StatusForbidden, r)
			c.Abort()
		}
		c.Set("user", user)
		c.Next()
	} else {
		fmt.Println(err)
		r.Message = "Authorization Failed"
		c.JSON(http.StatusForbidden, r)
		c.Abort()
	}
}
