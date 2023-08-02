package initializers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ilyasbabu/go-ecommerce/routes"
)

func SetupRoutes(ctx *gin.Engine) {
	routes.AdminRoutes(ctx)
	routes.UserRoutes(ctx)

	fmt.Println("--------------------AVAILABLE ROUTES--------------------")
	for _, routeInfo := range ctx.Routes() {
		fmt.Printf("%-6s %-20s %s\n", routeInfo.Method, routeInfo.Path, routeInfo.Handler)
	}
	fmt.Println("--------------------------------------------------------")
}
