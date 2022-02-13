package productController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"products-api/models"
	productService "products-api/services/product"
)

func Create(ctx *gin.Context) {
	var input models.CreateProductInput

	err := ctx.ShouldBindJSON(&input)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	product := productService.CreateProduct(input)

	ctx.JSON(200, product)
}

func FindMany(ctx *gin.Context) {
	products := productService.FindMany()

	ctx.JSON(http.StatusOK, products)
}

func FindOne(ctx *gin.Context) {
	productName, queryExists := ctx.GetQuery("productName")

	if !queryExists {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"error": "You should provide a product name"})
		return
	}

	product, productExists := productService.FindOneByName(productName)

	if !productExists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	ctx.JSON(http.StatusOK, product)

	return
}
