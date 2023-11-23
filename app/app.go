package app

import (
	"github.com/gin-gonic/gin"
	"github.com/go-bookstore-users-api/logger"
)

var (
	router = gin.Default()
)

func StartApplication() {
	logger.Info("Starting application...")
	mapUrls()
	router.Run()
}
