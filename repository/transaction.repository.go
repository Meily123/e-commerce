package repository

import (
	"WebAPI/config"
	"WebAPI/model"
)

type TransactionRepository interface {
	CreateTransaction(transaction model.Transaction) (model.Transaction, error)
	AppendProductTransaction(product model.TransactionProduct, transaction model.Transaction) (model.Transaction, error)
	ActivateTransaction(transaction model.Transaction) (model.Transaction, error)
	UpdateAddCalculationTransaction(transaction model.Transaction) (model.Transaction, error)
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
