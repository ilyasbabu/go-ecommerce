package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasbabu/go-ecommerce/controllers"
)

func AdminRoutes(ctx *gin.Engine) {
	admin := ctx.Group("admin")
	{
		admin.GET("ping/", controllers.Ping)
	}
}
