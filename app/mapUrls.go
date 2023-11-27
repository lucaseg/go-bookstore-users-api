package app

import (
	"github.com/go-bookstore-users-api/controller/ping"
	"github.com/go-bookstore-users-api/controller/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	// User urls
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
	router.PUT("/users/:user_id", users.UpdateUser)
	router.DELETE("/users/:user_id", users.DeleteUser)

	router.GET("/users/search", users.SearchUser)
	router.POST("/users/login", users.Login)
}
