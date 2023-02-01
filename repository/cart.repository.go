package repository

import (
	"WebAPI/config"
	"WebAPI/model"
)

type CartRepository interface {
	DeleteById(id string) error
	CreateCart(cart model.CartProduct, user model.User) (model.CartProduct, error)
	FindAll(user model.User) ([]model.CartProduct, error)
	FindById(id string) (model.CartProduct, error)
	FindByItemId(id string) (model.CartProduct, error)
	Update(id string, editRequest model.CartProductEditRequest) (model.CartProduct, error)
}

type cartRepository struct {
}

func NewCartRepository() *cartRepository {
	return &cartRepository{}
}

func (cartRepo *cartRepository) CreateCart(cart model.CartProduct, user model.User) (model.CartProduct, error) {
	// Create cart
	cart.UserId = user.Id

	err := config.ConnectToDatabase().Create(&cart).Error

	return cart, err
}

func (cartRepo *cartRepository) DeleteById(id string) error {
	// Delete by Id
	cart := model.CartProduct{}
	err := config.ConnectToDatabase().Delete(&cart, "id = ?", id).Error

	return err
}

func (cartRepo *cartRepository) FindAll(user model.User) ([]model.CartProduct, error) {
	// get user with all cart list
	err := config.ConnectToDatabase().Preload("Cart").First(&user).Error

	var carts []model.CartProduct

	productRepo := NewProductRepository()

	for _, cart := range user.Cart {
		product, _ := productRepo.FindById(cart.ItemId.String())
		cart.Item = product
		carts = append(carts, cart)
	}

	return carts, err
}

func (cartRepo *cartRepository) FindById(id string) (model.CartProduct, error) {
	// Find by id
	cart := model.CartProduct{}
	err := config.ConnectToDatabase().Find(&cart, "id = ?", id).Error

	if err != nil {
		return model.CartProduct{}, err
	}

	err = cart.EmptyProductStruct()

	return cart, err
}

func (cartRepo *cartRepository) FindByItemId(id string) (model.CartProduct, error) {
	// Find by id
	cart := model.CartProduct{}
	err := config.ConnectToDatabase().Find(&cart, "item_id = ?", id).Error

	if err != nil {
		return model.CartProduct{}, err
	}

	err = cart.EmptyProductStruct()

	return cart, err
}

func (cartRepo *cartRepository) Update(id string, editRequest model.CartProductEditRequest) (model.CartProduct, error) {
	// GET cart
	cart, err := cartRepo.FindById(id)

	if err != nil {
		return model.CartProduct{}, err
	}

	productRepo := NewProductRepository()
	product, err := productRepo.FindById(cart.ItemId.String())

	if err != nil {
		return model.CartProduct{}, err
	}

	cart.Item = product

	// Save changes
	err = config.ConnectToDatabase().Model(&cart).Select("Quantity").Updates(model.CartProduct{Quantity: editRequest.Quantity}).Error

	return cart, err
}
