package domain

import (
	"fmt"
	"strings"

	"github.com/go-bookstore-users-api/logger"
	"github.com/go-bookstore-users-api/repository/mysql/users_db"

	"github.com/go-bookstore-users-api/utils"
)

var (
	insertUserQuery   = "INSERT INTO users.users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	getUserQuery      = "SELECT * FROM users WHERE id = ?;"
	updateUserQuery   = "UPDATE users.users SET first_name = ?, last_name = ? WHERE id=?;"
	deleteUserQuery   = "DELETE FROM users WHERE id = ?;"
	findByStatusQuery = "SELECT * FROM users WHERE status=?;"
)

func (user *User) Get() *utils.RestError {

	stmt, err := users_db.Client.Prepare(getUserQuery)
	if err != nil {
		logger.Error("Error trying to get connection to database", err)
		return utils.InteralServerError("Error trying to get connection to database")
	}

	result := stmt.QueryRow(user.Id)
	if err := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password); err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return utils.NotFound("User not found")
		}
		return utils.InteralServerError("Error scanning results")
	}

	if result == nil {
		return utils.NotFound(fmt.Sprintf("user %d not found", user.Id))
	}

	return nil
}

func (user *User) Save() *utils.RestError {
	if err := user.Validate(); err != nil {
		return err
	}

	stmt, err := users_db.Client.Prepare(insertUserQuery)
	if err != nil {
		utils.InteralServerError("Error getting database connection")
	}

	defer stmt.Close()
	user.Password = utils.Encode(user.Password)
	resutl, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated, user.Status, user.Password)
	if err != nil {
		utils.InteralServerError("Error trying to insert the user into the data base")
	}

	fmt.Printf("%v", *user)

	userId, err := resutl.LastInsertId()

	if err != nil {
		utils.InteralServerError("Error trying to insert the user into the data base")
	}

	user.Id = userId
	return nil
}

func (user *User) Update() *utils.RestError {
	stmt, err := users_db.Client.Prepare(updateUserQuery)
	if err != nil {
		return utils.InteralServerError("Error getting database connection")
	}
	defer stmt.Close()

	_, errUpdate := stmt.Exec(user.FirstName, user.LastName)

	if errUpdate != nil {
		return utils.InteralServerError("Error updating user")
	}

	return nil
}

func (user *User) FindByStatus() ([]User, *utils.RestError) {
	stmt, err := users_db.Client.Prepare(findByStatusQuery)
	if err != nil {
		return nil, utils.InteralServerError("Error gettin database connection")
	}
	defer stmt.Close()

	rows, err := stmt.Query(user.Status)
	if err != nil {
		return nil, utils.InteralServerError("")
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status, &user.Password)
		if err != nil {
			return nil, utils.InteralServerError("")
		}
		results = append(results, user)
	}

	return results, nil
}

func (user *User) Delete() *utils.RestError {
	stmt, err := users_db.Client.Prepare(deleteUserQuery)
	if err != nil {
		return utils.InteralServerError("Error trying to get connection to database")
	}
	defer stmt.Close()

	_, deleteErr := stmt.Exec(user.Id)
	if deleteErr != nil {
		return utils.InteralServerError("Error executing query to delete user")
	}

	return nil
}
