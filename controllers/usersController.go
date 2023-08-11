package controllers

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ilyasbabu/go-ecommerce/models"
	"github.com/ilyasbabu/go-ecommerce/services"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Signup(c *gin.Context) {
	r := Response()
	email := c.PostForm("email")
	if email == "" {
		r.Message = "Email must be provided"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}

	password := c.PostForm("password")
	if password == "" {
		r.Message = "Password must be provided"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		r.Message = "Failed to hash password"
	}
	user := models.User{Email: email, Password: string(hash)}
	result := Db.Create(&user)
	if result.Error != nil {
		r.Message = result.Error.Error()
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			r.Message = "Cannot Signup! User with same email(" + user.Email + ") already exists"
		}
		c.JSON(http.StatusNotAcceptable, r)
		return
	}

	go services.SendMail(email, "Welcome to our store", "We are happy to have you onboard!")

	r.Status = "SUCCESS"
	r.Message = "Signed Up Successfully"
	c.JSON(http.StatusOK, r)
}

func Login(c *gin.Context) {
	r := Response()
	email := c.PostForm("email")
	if email == "" {
		r.Message = "Email must be provided"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}

	password := c.PostForm("password")
	if password == "" {
		r.Message = "Password must be provided"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}

	var user models.User
	err := Db.First(&user, "email=?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		r.Message = "Invalid Email or Password"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		r.Message = "Invalid Email or Password"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		r.Message = "Error generating token"
		c.JSON(http.StatusNotAcceptable, r)
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	r.Status = "SUCCESS"
	r.Message = "Logged In Successfully"
	r.Data = tokenString
	c.JSON(http.StatusOK, r)
}

func ValidateLogin(c *gin.Context) {
	r := Response()
	user, _ := c.Get("user")
	r.Status = "SUCCESS"
	r.Message = "Validated Successfully"
	r.Data = user
	c.JSON(http.StatusOK, r)
}
