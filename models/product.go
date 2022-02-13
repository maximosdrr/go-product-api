package models

type Product struct {
	Name        string
	Price       float32
	Description string
}

type CreateProductInput struct {
	Name        string `json:"name" binding:"required"`
	Price       float32   `json:"price" binding:"required"`
	Description string `json:"description" binding:"required"`
}
