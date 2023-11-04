package domain

import "encoding/json"

type PublicUserDto struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type PrivateUserDto struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

func (user *User) Marshal(isPublic bool) interface{} {
	userJson, _ := json.Marshal(user)

	if isPublic {
		var publicUser PublicUserDto
		json.Unmarshal(userJson, &publicUser)
		return publicUser
	}

	var privateUser PrivateUserDto
	json.Unmarshal(userJson, &privateUser)
	return privateUser
}
