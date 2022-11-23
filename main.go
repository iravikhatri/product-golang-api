package main

import (
	"github.com/gin-gonic/gin"

	"github.com/iravikhatri/go-api/controllers"
)

func main() {

	router := gin.Default()

	router.GET("/products", controllers.GetProducts)
	router.POST("/products", controllers.CreateProducts)
	router.GET("/product/:id", controllers.GetProduct)
	router.PATCH("/product/:id", controllers.UpdateProduct)
	router.DELETE("/product/:id", controllers.DeleteProduct)

	router.Run()
}
