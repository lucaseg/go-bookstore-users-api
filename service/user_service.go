package service

import (
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
	Login(email string, password string) (*domain.User, *utils.RestError)
}

func (s *userService) GetUser(userId int64) (*domain.User, *utils.RestError) {
	var userDto = domain.User{Id: userId}

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

func (s *userService) Login(login domain.Login) (*domain.User, *utils.RestError) {
	user := domain.User{
		Email:    login.Email,
		Password: login.Password,
	}

	if err := user.FindByEmailAndPassword(); err != nil {
		return nil, err
	}
	return &user, nil
}
