package repository

import (
	"WebAPI/config"
	"WebAPI/model"
)

type UserRepository interface {
	CreateUser(user model.User) (model.User, error)
	FindByUsername(loginRequest model.LoginRequest) (model.User, error)
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

	user := model.User{}
	err := config.ConnectToDatabase().Find(&user, "username = ?", loginRequest.Username).Error
	return user, err
}
