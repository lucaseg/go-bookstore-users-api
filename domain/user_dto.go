package domain

import (
	"strings"

	"github.com/go-bookstore-users-api/utils"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

var (
	minimumLenghtPassword = 8
)

func (user *User) Validate() *utils.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return utils.BadRequest("Invalid email.")
	}

	// Validate password
	password := strings.TrimSpace(user.Password)
	if password == "" || len(password) < minimumLenghtPassword {
		return utils.BadRequest("Invalid password")
	}
	return nil
}
