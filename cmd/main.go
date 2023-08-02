package main

import (
	"github.com/gin-gonic/gin"
	i "github.com/ilyasbabu/go-ecommerce/initializers"
)

func main() {
	r := gin.Default()
	i.ConnectDatabase()
	i.LoadEnv()
	i.SetupRoutes(r)
	r.Run(":8080")
}
