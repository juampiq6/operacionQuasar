package ginserver

import (
	"os"

	"github.com/gin-gonic/gin"
)

func GinStartServing() {
	router := gin.Default()
	router.POST("/topsecret", postTopSecretHandler)
	router.POST("/topsecret_split/:name", postTopSecretSplitHandler)
	router.GET("/topsecret_split/", getTopSecretSplitHandler)
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "80"
	}
	router.Run(":" + port)

}
