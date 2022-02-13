package productService

import (
	"products-api/models"
)

var products = make([]models.Product, 0)

func CreateProduct(product models.CreateProductInput) models.Product {
	productInstance := models.Product{
		Name: product.Name,
		Description: product.Description,
		Price: product.Price,
	}
	products = append(products, productInstance)

	return productInstance
}

func FindMany() []models.Product {
	return products
}

func FindOneByName(name string) (models.Product, bool) {
	for _, product := range products {
		if product.Name == name {
			return product, true
		}
	}

	return models.Product{
		Name:        "NOT FOUND",
		Price:       0.0,
		Description: "NOT FOUND",
	}, false
}
