package app

import (
	"bookstore_user-api/logger"

	"github.com/gin-gonic/gin"
)

var (
	r = gin.Default()
)

func StartApplication() {
	mapUrls()

	logger.Info("about to start the application...")
	r.Run(":8082")

}
