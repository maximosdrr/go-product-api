package products

import (
	"github.com/gin-gonic/gin"
	productController "products-api/controllers/product"
)

func GenerateProductsRoute(routerGroup *gin.RouterGroup) *gin.RouterGroup {
	products := routerGroup.Group("products")

	products.POST("/create", productController.Create)
	products.GET("/find", productController.FindMany)
	products.GET("find-one", productController.FindOne)

	return products
}
