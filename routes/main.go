package routes

import (
	"github.com/gin-gonic/gin"
	"products-api/routes/products"
)

func GenerateRoutes(server *gin.Engine) {
	main := server.Group("api/v1")
	products.GenerateProductsRoute(main)

}
