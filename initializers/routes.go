package initializers

import (
	"github.com/gin-gonic/gin"
	"github.com/ilyasbabu/go-ecommerce/routes"
)

func SetupRoutes(ctx *gin.Engine) {
	routes.AdminRoutes(ctx)
	routes.UserRoutes(ctx)
}
