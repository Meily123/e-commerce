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
	UpdateAdminById(id string) error
	Update(user model.User, editRequest model.UserEditRequest) (model.User, error)
}

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (userRepo *userRepository) CreateUser(user model.User) (model.User, error) {
	// Create user
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
	// Delete by Id
	user := model.User{}
	err := config.ConnectToDatabase().Delete(&user, "id = ?", id).Error

	return err
}

func (userRepo *userRepository) FindAll() ([]model.User, error) {
	// Find all
	var users []model.User
	err := config.ConnectToDatabase().Find(&users).Error

	return users, err
}

func (userRepo *userRepository) UpdateAdminById(id string) error {
	// Update to admin by id
	var user model.User
	err := config.ConnectToDatabase().Model(&user).Where("id = ?", id).Update("is_admin", "true").Error

	return err
}

func (userRepo *userRepository) Update(user model.User, editRequest model.UserEditRequest) (model.User, error) {
	// GET user
	config.ConnectToDatabase().First(&user)

	// Update user
	user.Name = editRequest.Name
	user.Username = editRequest.Username
	user.Password = editRequest.Password
	user.Address = editRequest.Address
	user.Email = editRequest.Address

	// Save changes
	err := config.ConnectToDatabase().Save(&user).Error

	return user, err
}
