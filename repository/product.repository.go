package repository

import (
	"WebAPI/config"
	"WebAPI/model"
)

type ProductRepository interface {
	CreateProduct(product model.Product) (model.Product, error)
}

type productRepository struct {
}

func NewProductRepository() *productRepository {
	return &productRepository{}
}

func (productRepo *productRepository) CreateProduct(product model.Product) (model.Product, error) {
	// Create product
	err := config.ConnectToDatabase().Create(&product).Error

	return product, err
}
