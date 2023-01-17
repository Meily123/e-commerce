package repository

import (
	"WebAPI/config"
	"WebAPI/model"
)

type ProductRepository interface {
	DeleteById(id string) error
	CreateProduct(product model.Product) (model.Product, error)
	FindAll() ([]model.Product, error)
	FindById(id string) (model.Product, error)
	Update(id string, editRequest model.ProductEditRequest) (model.Product, error)
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

func (productRepo *productRepository) DeleteById(id string) error {
	// Delete by Id
	product := model.Product{}
	err := config.ConnectToDatabase().Delete(&product, "id = ?", id).Error

	return err
}

func (productRepo *productRepository) FindAll() ([]model.Product, error) {
	// Find all
	var products []model.Product
	err := config.ConnectToDatabase().Find(&products).Error

	return products, err
}

func (productRepo *productRepository) FindById(id string) (model.Product, error) {
	// Find by id
	product := model.Product{}
	err := config.ConnectToDatabase().Find(&product, "id = ?", id).Error

	if err != nil {
		return model.Product{}, err
	}

	err = product.EmptyProductStruct()

	return product, err
}

func (productRepo *productRepository) Update(id string, editRequest model.ProductEditRequest) (model.Product, error) {
	// GET product
	product, err := productRepo.FindById(id)

	if err != nil {
		return model.Product{}, err
	}

	// Update product
	product.Name = editRequest.Name
	product.Stock = editRequest.Stock
	product.SellPrice = editRequest.SellPrice
	product.BasePrice = editRequest.BasePrice
	product.Descriptions = editRequest.Description

	// Save changes
	err = config.ConnectToDatabase().Save(&product).Error

	return product, err
}
