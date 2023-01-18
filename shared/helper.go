package shared

import (
	"WebAPI/config"
	"WebAPI/model"
	"WebAPI/repository"
)

func CompareAndPatchIfEmptyString(checkedValue string, patchValue string) string {
	if checkedValue == "" {
		return patchValue
	}
	return checkedValue
}

func GetCartsFromUser(u model.User) ([]model.CartProduct, error) {
	err := config.ConnectToDatabase().Preload("Cart").First(&u).Error

	if err != nil {
		return []model.CartProduct{}, err
	}

	var carts []model.CartProduct

	productRepo := repository.NewProductRepository()

	for _, cart := range u.Cart {
		product, _ := productRepo.FindById(cart.ItemId.String())
		cart.Item = product
		carts = append(carts, cart)
	}

	return carts, nil
}
