package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    uint64  `json:"id"`
	Name  string  `json:"name"`
	Price float32 `json:"price"`
}

var data = []Product{}

func GetProducts(ctx *gin.Context) {
	ctx.JSON(200, data)
}

func CreateProducts(ctx *gin.Context) {
	var product Product

	if err := ctx.BindJSON(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	data = append(data, product)
	ctx.JSON(http.StatusAccepted, &product)
}

func GetProduct(ctx *gin.Context) {
	id := ctx.Params.ByName("id")
	pid, _ := strconv.ParseUint(id, 0, 10)

	var product Product

	for idx, _ := range data {
		if data[idx].ID == pid {
			product = data[idx]
		}
	}

	if err := product; err.ID != 0 {
		ctx.AbortWithStatus(404)
		return
	}
	ctx.JSON(200, product)
}

func UpdateProduct(ctx *gin.Context) {
	// id := ctx.Params("id")

	var product Product
	if err := ctx.BindJSON(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusAccepted, &product)
}

func DeleteProduct(ctx *gin.Context) {
	// id := ctx.Params("id")

	var product Product
	if err := ctx.BindJSON(&product); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusAccepted, &product)
}
