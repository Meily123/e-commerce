package repository

import (
	"WebAPI/config"
	"WebAPI/model"
)

type UserRepository interface {
	CreateUser(user model.User) (model.User, error)
	FindAll() ([]model.User, error)
	FindByUsername(loginRequest model.LoginRequest) (model.User, error)
	FindById(id string) (model.User, error)
	DeleteById(id string) error
}

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (userRepo *userRepository) CreateUser(user model.User) (model.User, error) {
	err := config.ConnectToDatabase().Create(&user).Error

	return user, err
}

func (userRepo *userRepository) FindByUsername(loginRequest model.LoginRequest) (model.User, error) {
	// Find By Username
	user := model.User{}
	err := config.ConnectToDatabase().Find(&user, "username = ?", loginRequest.Username).Error

	return user, err
}

func (userRepo *userRepository) FindById(id string) (model.User, error) {
	// Find by id
	user := model.User{}
	err := config.ConnectToDatabase().Find(&user, "id = ?", id).Error

	return user, err
}

func (userRepo *userRepository) DeleteById(id string) error {
	// Find by id
	user := model.User{}
	err := config.ConnectToDatabase().Delete(&user, "id = ?", id).Error

	return err
}

func (userRepo *userRepository) FindAll() ([]model.User, error) {
	// Find by id
	var users []model.User
	err := config.ConnectToDatabase().Find(&users).Error

	return users, err
}
