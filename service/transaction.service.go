package service

import (
	"WebAPI/model"
	"WebAPI/repository"
	"WebAPI/shared"
	"errors"
	"fmt"
)

type TransactionService interface {
	DetailTransaction(user model.User) (model.Transaction, error)
	CreateTransaction(user model.User) (model.Transaction, error)
	VerifyPaymentTransaction(id string) error
	AllTransaction() ([]model.Transaction, error)
	FindByIdTransaction(id string, user model.User) (model.Transaction, error)
	SelfAllTransaction(user model.User) ([]model.Transaction, error)
	SummaryTransaction() (model.TransactionSummaryResponse, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewTransactionService(transactionRepo repository.TransactionRepository) *transactionService {
	return &transactionService{transactionRepo}
}

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

	if len(carts) == 0 {
		return model.Transaction{}, errors.New("the cart must not be empty")
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

		sufficient, _ := checkSufficient(cart)

		if sufficient {
			transaction, err = transactionServ.transactionRepository.AppendProductTransaction(transactionProduct, transaction)
			if err != nil {
				return model.Transaction{}, err
			}

			totalItem++
			totalMargin += transactionProduct.Margin
			totalSum += transactionProduct.Sum

			productRepo := repository.NewProductRepository()
			productServ := NewProductService(productRepo)
			product, _ := productRepo.FindById(cart.ItemId.String())
			_ = productServ.UpdateStockProduct(cart.ItemId.String(), product.Stock-cart.Quantity)
			fmt.Println(cart.ItemId.String())
		}

		cartRepo := repository.NewCartRepository()
		_ = cartRepo.DeleteById(cart.Id.String())
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

	var transactionProducts []model.TransactionProduct

	for _, cart := range carts {
		transactionProduct := model.TransactionProduct{
			Item:     cart.Item,
			ItemId:   cart.ItemId,
			Quantity: cart.Quantity,
			Sum:      cart.Item.SellPrice * cart.Quantity,
			Margin:   (cart.Item.SellPrice - cart.Item.BasePrice) * cart.Quantity,
		}

		sufficient, _ := checkSufficient(cart)
		if sufficient {
			transactionProducts = append(transactionProducts, transactionProduct)

			totalItem++
			totalMargin += transactionProduct.Margin
			totalSum += transactionProduct.Sum
		}
	}

	transaction.Products = transactionProducts
	transaction.TotalMargin = totalMargin
	transaction.TotalItem = totalItem
	transaction.TotalSum = totalSum

	return transaction, err
}

func (transactionServ *transactionService) VerifyPaymentTransaction(id string) error {
	transaction, err := transactionServ.transactionRepository.FindByIdTransaction(id)
	if err != nil {
		return err
	}

	err = transactionServ.transactionRepository.VerifyPaymentTransaction(transaction)
	if err != nil {
		return err
	}

	return err
}

func (transactionServ *transactionService) AllTransaction() ([]model.Transaction, error) {

	transactions, err := transactionServ.transactionRepository.FindAllTransaction()
	if err != nil {
		return []model.Transaction{}, err
	}

	var loadedTransactions []model.Transaction

	// load the transaction product
	for _, transaction := range transactions {
		transaction.Products, _ = transactionServ.transactionRepository.FindTransactionProducts(transaction.Id.String())
		var productItems []model.TransactionProduct
		for _, transactionProduct := range transaction.Products {
			productRepo := repository.NewProductRepository()
			transactionProduct.Item, _ = productRepo.FindById(transactionProduct.ItemId.String())
			productItems = append(productItems, transactionProduct)
		}
		transaction.Products = productItems
		loadedTransactions = append(loadedTransactions, transaction)
	}

	return loadedTransactions, err
}

func (transactionServ *transactionService) FindByIdTransaction(id string, user model.User) (model.Transaction, error) {

	transaction, err := transactionServ.transactionRepository.FindByIdTransaction(id)
	if err != nil {
		return model.Transaction{}, err
	}

	if user.IsAdmin == false {
		// check if its user's
		if transaction.UserId != user.Id {
			return model.Transaction{}, err
		}
	}

	return transaction, err
}

func (transactionServ *transactionService) SelfAllTransaction(user model.User) ([]model.Transaction, error) {

	transactions, err := transactionServ.transactionRepository.SelfFindAllTransaction(user)
	if err != nil {
		return []model.Transaction{}, err
	}

	// load the transaction product
	for _, transaction := range transactions {
		transaction.Products, _ = transactionServ.transactionRepository.FindTransactionProducts(transaction.Id.String())
	}

	return transactions, err
}

func (transactionServ *transactionService) SummaryTransaction() (model.TransactionSummaryResponse, error) {
	var transactionSummary model.TransactionSummaryResponse

	transactions, err := transactionServ.AllTransaction()
	if err != nil {
		return transactionSummary, err
	}

	for _, transaction := range transactions {
		if transaction.IsPaid == true {
			transactionSummary.TotalMarginSoldProduct += transaction.TotalMargin
			transactionSummary.TotalSoldProduct += transaction.TotalItem
			transactionSummary.TotalSumSoldProduct += transaction.TotalSum
		}
	}

	return transactionSummary, err
}

func checkSufficient(cart model.CartProduct) (bool, error) {
	productRepo := repository.NewProductRepository()
	product, err := productRepo.FindById(cart.ItemId.String())

	if err != nil {
		return false, errors.New("error while getting product")
	}

	if product.Stock > cart.Quantity {
		return true, nil
	}

	return false, nil
}
