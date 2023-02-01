package service

import (
	"WebAPI/model"
	"WebAPI/repository"
	"errors"
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

	// get product
	productRepo := repository.NewProductRepository()
	product, err := productRepo.FindById(cartRequest.ItemId)
	if err != nil {
		return model.CartProduct{}, err
	}

	// find product in cart
	existCartProduct, err := cartServ.cartRepository.FindByItemId(cartRequest.ItemId)

	if existCartProduct.Id.String() != "00000000-0000-0000-0000-000000000000" {
		return model.CartProduct{}, errors.New("cart product already exist")
	}

	//check if product stock sufficient
	if product.Stock < cartRequest.Quantity {
		return model.CartProduct{}, errors.New("insufficient product stock")
	}

	// create cart struct
	cart := model.CartProduct{
		ItemId:   product.Id,
		Quantity: cartRequest.Quantity,
	}

	// create cart
	cart, err = cartServ.cartRepository.CreateCart(cart, user)
	if err != nil {
		return model.CartProduct{}, err
	}

	cart.Item = product

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

	for _, cart := range carts {
		if cart.Item.Stock < cart.Quantity {
			cart.Quantity = cart.Item.Stock

			editRequest := model.CartProductEditRequest{
				Quantity: cart.Item.Stock,
			}
			_, err = cartServ.cartRepository.Update(cart.Id.String(), editRequest)

			if err != nil {
				return []model.CartProduct{}, err
			}
		}
	}

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

	// get product
	productRepo := repository.NewProductRepository()
	product, err := productRepo.FindById(id)
	if err != nil {
		return model.CartProduct{}, err
	}

	// check if stock sufficient
	if product.Stock < editRequest.Quantity {
		return model.CartProduct{}, errors.New("insufficient product stock")
	}

	// update cart
	cart, err := cartServ.cartRepository.Update(id, editRequest)
	if err != nil {
		return model.CartProduct{}, err
	}

	return cart, err
}

func (cartServ *cartService) checkUserCart(id string, userId string) error {
	//find the cart
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
