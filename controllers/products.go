package controllers

import "github.com/gin-gonic/gin"

func Ping(ctx *gin.Context) {
	ctx.String(200, "Pong")
}

func Ling(ctx *gin.Context) {
	ctx.String(200, "Long")
}
