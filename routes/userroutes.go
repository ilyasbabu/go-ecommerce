package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasbabu/go-ecommerce/controllers"
	"github.com/ilyasbabu/go-ecommerce/middlewares"
)

func UserRoutes(ctx *gin.Engine) {
	user := ctx.Group("/")
	{
		user.GET("/ping", controllers.Ping)
		user.POST("/signup", controllers.Signup)
		user.POST("/login", controllers.Login)
		user.POST("/validate", middlewares.RequireAuth, controllers.ValidateLogin)
	}
}
