package service

import (
	"errors"
	"myapp-backend/controller/request"
	"myapp-backend/domain"
	"myapp-backend/repository"
	"myapp-backend/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (userService *UserService) CreateUser(user request.User) error {
	userPointer, err := userService.userRepository.FindByEmail(user.Email)

	if err != nil {
		return err
	}

	if userPointer != nil {
		return errors.New("email is already in use")
	}

	hashedPassword, hashError := HashPassword(user.Password)

	if hashError != nil {
		return errors.New("error occured whilte password is hashing")
	}

	return userService.userRepository.CreateUser(domain.User{
		Name:           user.Name,
		Surname:        user.Surname,
		Email:          user.Email,
		HashedPassword: hashedPassword,
	})
}

func (userService *UserService) Authenticate(email string, password string) (string, error) {
	user, err := userService.userRepository.FindByEmail(email)

	if err != nil {
		return "", err
	}

	if user == nil {
		return "", errors.New("user not found")
	}

	if !CheckPassword(user.HashedPassword, password) {
		return "", errors.New("the password is incorrect")
	}

	token, err := utils.GenerateJWT(user.Email)
	if err != nil {
		return "", errors.New("error occured while token is generating")
	}

	return token, nil
}

func CheckPassword(hashedPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
