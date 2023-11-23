package service

import (
	"fmt"

	"github.com/go-bookstore-users-api/domain"
	"github.com/go-bookstore-users-api/utils"
)

var (
	UserService userService = userService{}
)

type userService struct {
}

type userServiceInterface interface {
	GetUser(userId int64) (*domain.User, *utils.RestError)
	CreateUser(user domain.User) (*domain.User, *utils.RestError)
	UpdateUser(user domain.User) (*domain.User, *utils.RestError)
	DeleteUser(userId int64) *utils.RestError
}

func (s *userService) GetUser(userId int64) (*domain.User, *utils.RestError) {
	var userDto = domain.User{Id: userId}

	fmt.Printf("Valor de la variable %v \n", userDto)
	fmt.Printf("Dirección de memoria de userDto: %p\n", &userDto)

	var ptr *domain.User = &userDto
	fmt.Printf("Valor del puntero %p \n", ptr)
	fmt.Printf("Direccion de memoria del puntero %p \n", &ptr)

	err := userDto.Get()
	if err != nil {
		return nil, err
	}

	return &userDto, nil
}

func (s *userService) CreateUser(user domain.User) (*domain.User, *utils.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}
	user.DateCreated = utils.GetDateNow()
	saveError := user.Save()
	if saveError != nil {
		return nil, saveError
	}

	return &user, nil
}

func (s *userService) UpdateUser(user domain.User) (*domain.User, *utils.RestError) {
	var err *utils.RestError

	originalUser, err := s.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	originalUser.FirstName = user.FirstName
	originalUser.LastName = user.LastName
	originalUser.Email = user.Email

	err = originalUser.Update()
	if err != nil {
		return nil, err
	}

	return originalUser, nil
}

func (s *userService) DeleteUser(userId int64) *utils.RestError {
	var user domain.User
	user.Id = userId

	if err := user.Delete(); err != nil {
		return err
	}
	return nil
}
