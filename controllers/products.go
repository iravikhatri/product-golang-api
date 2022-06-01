package controllers

import "github.com/gin-gonic/gin"

func GetProducts(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Welcome",
	})
}
