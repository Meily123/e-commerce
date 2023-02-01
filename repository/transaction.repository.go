package repository

import (
	"WebAPI/config"
	"WebAPI/model"
	"time"
)

type TransactionRepository interface {
	CreateTransaction(transaction model.Transaction) (model.Transaction, error)
	AppendProductTransaction(product model.TransactionProduct, transaction model.Transaction) (model.Transaction, error)
	ActivateTransaction(transaction model.Transaction) (model.Transaction, error)
	UpdateAddCalculationTransaction(transaction model.Transaction) (model.Transaction, error)
	VerifyPaymentTransaction(transaction model.Transaction) error
	FindByIdTransaction(id string) (model.Transaction, error)
	FindAllTransaction() ([]model.Transaction, error)
	SelfFindAllTransaction(user model.User) ([]model.Transaction, error)
	DeleteTransaction(id string) error
	FindTransactionProducts(id string) ([]model.TransactionProduct, error)
}

type transactionRepository struct {
}

func NewTransactionRepository() *transactionRepository {
	return &transactionRepository{}
}

func (transactionRepo *transactionRepository) CreateTransaction(transaction model.Transaction) (model.Transaction, error) {
	// Create transaction
	err := config.ConnectToDatabase().Create(&transaction).Error

	return transaction, err
}

func (transactionRepo *transactionRepository) AppendProductTransaction(product model.TransactionProduct, transaction model.Transaction) (model.Transaction, error) {
	err := config.ConnectToDatabase().Model(&transaction).Association("Products").Append(&product)

	return transaction, err
}

func (transactionRepo *transactionRepository) DeleteTransaction(id string) error {
	// Delete by Id
	transaction := model.Transaction{}
	err := config.ConnectToDatabase().Delete(&transaction, "id = ?", id).Error

	return err
}

func (transactionRepo *transactionRepository) UpdateAddCalculationTransaction(transaction model.Transaction) (model.Transaction, error) {
	err := config.ConnectToDatabase().Model(&transaction).Select("TotalSum", "TotalMargin", "TotalItem").Updates(transaction).Error

	return transaction, err
}

func (transactionRepo *transactionRepository) ActivateTransaction(transaction model.Transaction) (model.Transaction, error) {
	err := config.ConnectToDatabase().Model(&transaction).Select("IsActive").Updates(model.Transaction{IsActive: true}).Error

	return transaction, err
}

func (transactionRepo *transactionRepository) VerifyPaymentTransaction(transaction model.Transaction) error {
	err := config.ConnectToDatabase().Model(&transaction).Select("IsPaid", "PaidAt").Updates(model.Transaction{IsPaid: true, PaidAt: time.Now().UnixNano()}).Error

	return err
}

func (transactionRepo *transactionRepository) FindByIdTransaction(id string) (model.Transaction, error) {
	// Find by id
	transaction := model.Transaction{}
	err := config.ConnectToDatabase().Find(&transaction, "id = ?", id).Error

	if err != nil {
		return model.Transaction{}, err
	}

	return transaction, err
}

func (transactionRepo *transactionRepository) FindAllTransaction() ([]model.Transaction, error) {
	// Find by id
	var transaction []model.Transaction
	err := config.ConnectToDatabase().Find(&transaction).Error

	if err != nil {
		return []model.Transaction{}, err
	}

	return transaction, err
}

func (transactionRepo *transactionRepository) SelfFindAllTransaction(user model.User) ([]model.Transaction, error) {
	// Find by id
	var transaction []model.Transaction
	err := config.ConnectToDatabase().Find(&transaction).Where("UserId = ", user.Id).Error

	if err != nil {
		return []model.Transaction{}, err
	}

	return transaction, err
}

func (transactionRepo *transactionRepository) FindTransactionProducts(id string) ([]model.TransactionProduct, error) {
	// find transaction products by id transaction
	var products []model.TransactionProduct
	err := config.ConnectToDatabase().Find(&products).Where("transaction_id = ", id).Error

	if err != nil {
		return []model.TransactionProduct{}, err
	}

	return products, err
}
