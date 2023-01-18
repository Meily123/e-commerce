package service

import (
	"WebAPI/model"
	"WebAPI/repository"
	"WebAPI/shared"
)

type TransactionService interface {
	DetailTransaction(user model.User) (model.Transaction, error)
	CreateTransaction(user model.User) (model.Transaction, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) *transactionService {
	return &transactionService{transactionRepo}
}

// Create transaction
// Get all transaction's transaction item make sure its at least 1 product here
// create transaction, add transaction item sambil hitung, hitung jumlah dan total, activate, hapus transaction yang ada

func (transactionServ *transactionService) CreateTransaction(user model.User) (model.Transaction, error) {

	// create transaction base
	transaction := model.Transaction{
		UserId: user.Id,
	}
	transaction, err := transactionServ.transactionRepository.CreateTransaction(transaction)
	if err != nil {
		return model.Transaction{}, err
	}

	carts, err := shared.GetCartsFromUser(user)
	if err != nil {
		return model.Transaction{}, err
	}

	// loop through carts and add cart product into transaction product
	// count total item, sum, and margin transaction
	totalItem := 0
	totalSum := 0
	totalMargin := 0

	for _, cart := range carts {
		transactionProduct := model.TransactionProduct{
			Item:     cart.Item,
			ItemId:   cart.ItemId,
			Quantity: cart.Quantity,
			Sum:      cart.Item.SellPrice * cart.Quantity,
			Margin:   (cart.Item.SellPrice - cart.Item.BasePrice) * cart.Quantity,
		}

		transaction, err = transactionServ.transactionRepository.AppendProductTransaction(transactionProduct, transaction)
		if err != nil {
			return model.Transaction{}, err
		}

		totalItem++
		totalMargin += transactionProduct.Margin
		totalSum += transactionProduct.Sum
	}

	transaction.TotalMargin = totalMargin
	transaction.TotalItem = totalItem
	transaction.TotalSum = totalSum

	transaction, err = transactionServ.transactionRepository.UpdateAddCalculationTransaction(transaction)
	if err != nil {
		return model.Transaction{}, err
	}

	transaction, err = transactionServ.transactionRepository.ActivateTransaction(transaction)
	if err != nil {
		return model.Transaction{}, err
	}

	// delete cart product from cart
	cartRepo := repository.NewCartRepository()

	for _, cart := range carts {
		_ = cartRepo.DeleteById(cart.Id.String())
	}

	return transaction, err
}

func (transactionServ *transactionService) DetailTransaction(user model.User) (model.Transaction, error) {
	// create transaction base
	transaction := model.Transaction{
		UserId: user.Id,
	}

	carts, err := shared.GetCartsFromUser(user)
	if err != nil {
		return model.Transaction{}, err
	}

	// loop through carts and add cart product into transaction product
	// count total item, sum, and margin transaction
	totalItem := 0
	totalSum := 0
	totalMargin := 0

	transactionProducts := []model.TransactionProduct{}

	for _, cart := range carts {
		transactionProduct := model.TransactionProduct{
			Item:     cart.Item,
			ItemId:   cart.ItemId,
			Quantity: cart.Quantity,
			Sum:      cart.Item.SellPrice * cart.Quantity,
			Margin:   (cart.Item.SellPrice - cart.Item.BasePrice) * cart.Quantity,
		}

		transactionProducts = append(transactionProducts, transactionProduct)

		totalItem++
		totalMargin += transactionProduct.Margin
		totalSum += transactionProduct.Sum
	}

	transaction.Products = transactionProducts
	transaction.TotalMargin = totalMargin
	transaction.TotalItem = totalItem
	transaction.TotalSum = totalSum

	return transaction, err
}