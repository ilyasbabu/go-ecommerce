package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(200, "Pong")
	})
	r.Run(":8080")
}
