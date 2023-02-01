package service

import (
	"WebAPI/model"
	"WebAPI/repository"
)

type ProductService interface {
	CreateProduct(productRequest model.ProductRequest) (model.Product, error)
	FindById(id string) (model.Product, error)
	DeleteById(id string) error
	FindAll() ([]model.Product, error)
	GetFindById(id string) (model.Product, error)
	UpdateProduct(id string, editRequest model.ProductEditRequest) (model.Product, error)
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
	newProduct, err := productServ.productRepository.CreateProduct(product)

	return newProduct, err
}

func (productServ *productService) FindById(id string) (model.Product, error) {
	// find product by id
	product, err := productServ.productRepository.FindById(id)

	return product, err
}

func (productServ *productService) DeleteById(id string) error {
	// find the product
	_, err := productServ.productRepository.FindById(id)

	if err != nil {
		return err
	}

	// delete product by id
	err = productServ.productRepository.DeleteById(id)

	return err
}

func (productServ *productService) FindAll() ([]model.Product, error) {
	// find all product
	products, err := productServ.productRepository.FindAll()

	return products, err
}

func (productServ *productService) GetFindById(id string) (model.Product, error) {
	// find all product
	product, err := productServ.productRepository.FindById(id)

	return product, err
}

func (productServ *productService) UpdateProduct(id string, editRequest model.ProductEditRequest) (model.Product, error) {
	// update product
	product, err := productServ.productRepository.Update(id, editRequest)

	return product, err
}

func (productServ *productService) UpdateStockProduct(id string, stock int) error {
	// update stock product
	product, err := productServ.productRepository.FindById(id)
	err = productServ.productRepository.UpdateStock(stock, product)

	return err
}
