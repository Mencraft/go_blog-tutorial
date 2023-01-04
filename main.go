package main

import (
	"blog/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default() //new gin router initialization
	router.GET("/", func(context *gin.Context) {
		infrastructure.LoadEnv()
		infrastructure.NewDatabase()
		context.JSON(http.StatusOK, gin.H{"data": "Hello World !"})
	}) // first endpoint returns Hello World
	router.Run(":8000") //running application, Default port is 8080
}
