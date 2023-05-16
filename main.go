package main

import (
	"jjcaptain/receipt-processor/handlers"

	"github.com/gin-gonic/gin"
)

// Entry point of the program. Maps the endpoints and starts the web server.
func main() {
	router := gin.Default()

	receipts := router.Group("/receipts")
	{
		receipts.POST("/process", handlers.ProcessReceipt)
		receipts.GET("/:id/points", handlers.GetPointsForReceipt)	
	}

	router.Run("localhost:8080")
}
