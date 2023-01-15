package service

import (
	"WebAPI/model"
	"WebAPI/repository"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"os"
	"time"
)

type UserService interface {
	CreateUser(user model.UserRequest) (model.User, error)
	LoginUser(loginRequest model.LoginRequest) (string, error)
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{userRepo}
}

func ByteToString(byteVar []byte) string {
	return fmt.Sprintf("%s", byteVar)
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
		Password: ByteToString(hash),
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
