package service

import (
	"WebAPI/model"
	"WebAPI/repository"
	"WebAPI/shared"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type UserService interface {
	CreateUser(user model.UserRequest) (model.User, error)
	LoginUser(loginRequest model.LoginRequest) (string, error)
	FindById(id string) (model.User, error)
	DeleteById(selfRequestUser model.User) error
	FindAll() ([]model.User, error)
	GetFindById(id string) (model.User, error)
	UpdateAdminById(id string) (model.User, error)
	UpdateUser(id string, editRequest model.UserEditRequest) (model.User, error)
	SelfUpdateUser(selfUserRequest model.User, editRequest model.UserEditRequest) (model.User, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{userRepo}
}

func (userServ *userService) CreateUser(userRequest model.UserRequest) (model.User, error) {
	// hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(userRequest.Password), 12)

	if err != nil {
		errorMassage := errors.New("fail to hash the password")
		return model.User{}, errorMassage
	}

	// userRequest to user
	user := model.User{
		Name:     userRequest.Name,
		Username: userRequest.Username,
		Address:  userRequest.Address,
		Password: string(hash),
		Email:    userRequest.Email,
	}

	// point to repository CreateUser
	newUser, err := userServ.userRepository.CreateUser(user)
	return newUser, err
}

func (userServ *userService) LoginUser(loginRequest model.LoginRequest) (string, error) {
	// find user by username
	user, err := userServ.userRepository.FindByUsername(loginRequest)

	if err != nil {
		return "", err
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequest.Password))
	if err != nil {
		return "", err
	}

	// generate jwt token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.Id,
		"exp": time.Now().Add(time.Hour * 2).UnixNano(),
	})

	// Sign and get complete encode token
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return tokenString, err
	}

	return tokenString, nil
}

func (userServ *userService) FindById(id string) (model.User, error) {
	// find user by id
	user, err := userServ.userRepository.FindById(id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (userServ *userService) DeleteById(selfRequestUser model.User) error {
	// delete user by id
	err := userServ.userRepository.DeleteById(selfRequestUser.Id.String())

	if err != nil {
		return err
	}

	return nil
}

func (userServ *userService) FindAll() ([]model.User, error) {
	// find all user
	users, err := userServ.userRepository.FindAll()

	if err != nil {
		return []model.User{}, err
	}

	return users, nil
}

func (userServ *userService) GetFindById(id string) (model.User, error) {
	// find all user
	user, err := userServ.userRepository.FindById(id)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (userServ *userService) UpdateAdminById(id string) (model.User, error) {
	// find all user
	err := userServ.userRepository.UpdateAdminById(id)

	if err != nil {
		return model.User{}, err
	}

	user, _ := userServ.FindById(id)

	return user, nil
}

func (userServ *userService) UpdateUser(id string, editRequest model.UserEditRequest) (model.User, error) {

	// Find user
	user, err := userServ.GetFindById(id)

	if err != nil {
		return model.User{}, err
	}

	user, err = userServ.SelfUpdateUser(user, editRequest)
	return user, err
}

func (userServ *userService) SelfUpdateUser(selfUserRequest model.User, editRequest model.UserEditRequest) (model.User, error) {

	user := selfUserRequest

	// check edit request value if none replace with old value
	editRequest.Name = shared.CompareAndPatchIfEmptyString(editRequest.Name, user.Name)
	editRequest.Username = shared.CompareAndPatchIfEmptyString(editRequest.Username, user.Username)
	editRequest.Email = shared.CompareAndPatchIfEmptyString(editRequest.Email, user.Email)
	editRequest.Address = shared.CompareAndPatchIfEmptyString(editRequest.Address, user.Address)
	editRequest.Password = shared.CompareAndPatchIfEmptyString(editRequest.Password, user.Password)

	// check if it is the old password or new one
	if user.Password != editRequest.Password {
		// hash the password
		hash, err := bcrypt.GenerateFromPassword([]byte(editRequest.Password), 12)

		editRequest.Password = string(hash)

		if err != nil {
			errorMassage := errors.New("fail to hash the password")
			return model.User{}, errorMassage
		}
	}

	// update user
	user, err := userServ.userRepository.Update(user, editRequest)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
