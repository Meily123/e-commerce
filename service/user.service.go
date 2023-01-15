package service

import (
	"WebAPI/model"
	"WebAPI/repository"
)

type UserService interface {
	CreateUser(user model.UserRequest) (model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{userRepo}
}

func (userServ *userService) CreateUser(userRequest model.UserRequest) (model.User, error) {
	// process the password

	// userRequest to user
	user := model.User{
		Name:     userRequest.Name,
		Username: userRequest.UserName,
		Address:  userRequest.Address,
		Password: userRequest.Password,
		Email:    userRequest.Email,
	}

	newUser, err := userServ.userRepository.CreateUser(user)
	return newUser, err
}
