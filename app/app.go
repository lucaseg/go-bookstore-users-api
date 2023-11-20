package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-bookstore-users-api/repository/mysql/users_db"
)

var (
	router = gin.Default()
)

func StartApplication() {
	users_db.Init()
	mapUrls()
	router.Run()
}
