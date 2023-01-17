package service

import (
	"WebAPI/config"
	"WebAPI/model"
	"WebAPI/repository"
)

type CartService interface {
	CreateCart(cartRequest model.CartProductRequest, user model.User) (model.CartProduct, error)
	DeleteById(id string, user model.User) error
	FindAll(user model.User) ([]model.CartProduct, error)
	UpdateCart(id string, editRequest model.CartProductEditRequest, user model.User) (model.CartProduct, error)
}

type cartService struct {
	cartRepository repository.CartRepository
}

func NewCartService(cartRepo repository.CartRepository) *cartService {
	return &cartService{cartRepo}
}

func (cartServ *cartService) CreateCart(cartRequest model.CartProductRequest, user model.User) (model.CartProduct, error) {

	// cartRequest to cart
	productRepo := repository.NewProductRepository()
	product, err := productRepo.FindById(cartRequest.ItemId)
	cart := model.CartProduct{
		ItemId:   product.Id,
		Item:     product,
		Quantity: cartRequest.Quantity,
	}

	// get user with all cart list
	err = config.ConnectToDatabase().Preload("Cart").First(&user).Error

	// point to repository CreateCart
	cart, err = cartServ.cartRepository.CreateCart(cart, user)

	return cart, err
}

func (cartServ *cartService) DeleteById(id string, user model.User) error {
	// check
	err := cartServ.checkUserCart(id, user.Id.String())
	if err != nil {
		return err
	}

	// delete cart by id
	err = cartServ.cartRepository.DeleteById(id)
	if err != nil {
		return err
	}

	return nil
}

func (cartServ *cartService) FindAll(user model.User) ([]model.CartProduct, error) {
	// find all cart
	carts, err := cartServ.cartRepository.FindAll(user)

	if err != nil {
		return []model.CartProduct{}, err
	}

	return carts, nil
}

func (cartServ *cartService) GetFindById(id string, user model.User) (model.CartProduct, error) {
	// check
	err := cartServ.checkUserCart(id, user.Id.String())
	if err != nil {
		return model.CartProduct{}, err
	}

	// find all cart
	cart, err := cartServ.cartRepository.FindById(id)

	if err != nil {
		return model.CartProduct{}, err
	}

	return cart, nil
}

func (cartServ *cartService) UpdateCart(id string, editRequest model.CartProductEditRequest, user model.User) (model.CartProduct, error) {
	// check
	err := cartServ.checkUserCart(id, user.Id.String())
	if err != nil {
		return model.CartProduct{}, err
	}

	// update cart
	cart, err := cartServ.cartRepository.Update(id, editRequest)
	if err != nil {
		return model.CartProduct{}, err
	}

	return cart, err
}

func (cartServ *cartService) checkUserCart(id string, userId string) error {
	//find he cart
	cart, err := cartServ.cartRepository.FindById(id)
	if err != nil {
		return err
	}

	// is this the user cart ?
	err = cart.IsThisUserCart(userId)
	if err != nil {
		return err
	}

	return nil
}
