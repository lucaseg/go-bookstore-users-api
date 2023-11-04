package service

import (
	"fmt"

	"github.com/go-bookstore-users-api/domain"
	"github.com/go-bookstore-users-api/utils"
)

func GetUser(userId int64) (*domain.User, *utils.RestError) {
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

func CreateUser(user domain.User) (*domain.User, *utils.RestError) {
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

func UpdateUser(user domain.User) (*domain.User, *utils.RestError) {
	var err *utils.RestError

	originalUser, err := GetUser(user.Id)
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

func DeleteUser(userId int64) *utils.RestError {
	var user domain.User
	user.Id = userId

	if err := user.Delete(); err != nil {
		return err
	}
	return nil
}
