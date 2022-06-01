package main

import (
	"github.com/gin-gonic/gin"

	"github.com/iravikhatri/go-api/controllers"
)

func main() {

	router := gin.Default()

	router.GET("/", controllers.GetProducts)

	router.Run()
}
