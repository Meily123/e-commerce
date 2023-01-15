package repository

import (
	"WebAPI/config"
	"WebAPI/model"
)

type UserRepository interface {
	CreateUser(user model.User) (model.User, error)
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
