package service

import (
	"WebAPI/model"
	"WebAPI/repository"
)

type ProductService interface {
	CreateProduct(productRequest model.ProductRequest) (model.Product, error)
}

type productService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) *productService {
	return &productService{productRepo}
}

func (productServ *productService) CreateProduct(productRequest model.ProductRequest) (model.Product, error) {

	// productRequest to product
	product := model.Product{
		Name:         productRequest.Name,
		Descriptions: productRequest.Description,
		Stock:        productRequest.Stock,
		SellPrice:    productRequest.SellPrice,
		BasePrice:    productRequest.BasePrice,
	}

	// point to repository CreateProduct
	newUser, err := productServ.productRepository.CreateProduct(product)
	return newUser, err
}
