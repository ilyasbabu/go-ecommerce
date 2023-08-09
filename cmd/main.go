package main

import (
	"github.com/gin-gonic/gin"
	i "github.com/ilyasbabu/go-ecommerce/initializers"
)

func init() {
	i.LoadEnv()
	i.ConnectDatabase()
	i.SyncDatabase()
}

func main() {
	r := gin.Default()
	i.SetupRoutes(r)
	r.Run(":8080")
}
