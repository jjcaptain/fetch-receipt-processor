package main

import (
	"jjcaptain/receipt-processor/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// https://go.dev/doc/tutorial/web-service-gin
	receipts := router.Group("/receipts")
	{
		receipts.POST("/process", handlers.ProcessReceipt)
		receipts.GET("/:id/points", handlers.GetPointsForReceipt)	
	}

	router.Run("localhost:8080")
}
