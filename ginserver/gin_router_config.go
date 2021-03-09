package ginserver

import (
	"github.com/gin-gonic/gin"
)

func GinStartServing() {
	router := gin.Default()
	router.POST("/topsecret", postTopSecretHandler)
	router.POST("/topsecret_split/:name", postTopSecretSplitHandler)
	router.GET("/topsecret_split/", getTopSecretSplitHandler)

	router.Run(":8080")

}
