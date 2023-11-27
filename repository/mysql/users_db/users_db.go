package users_db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {
	dataSoruceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "example", "127.0.0.1:3306", "users")

	var err error
	Client, err = sql.Open("mysql", dataSoruceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connection to database successfully")
}
