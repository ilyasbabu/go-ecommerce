package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasbabu/go-ecommerce/controllers"
)

func UserRoutes(ctx *gin.Engine) {
	user := ctx.Group("/")
	{
		user.GET("/ping", controllers.Ping)
	}
}
