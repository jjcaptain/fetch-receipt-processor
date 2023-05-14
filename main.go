package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", hello)

	router.Run("localhost:8080")
}

func hello(context *gin.Context) {
	context.JSON(http.StatusOK, "Hello World")
}
