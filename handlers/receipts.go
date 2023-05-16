package handlers

import (
	"jjcaptain/receipt-processor/data"
	"jjcaptain/receipt-processor/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type receiptId struct {
	ID string `json:"id"`
}

type receiptPoints struct {
	Points int `json:"points"`
}

// Process a single receipt. Generate an ID, calculate and store the amount of points the receipt
// is worth, and return the generated ID.
func ProcessReceipt(context *gin.Context) {
	var newReceipt types.Receipt

	if err := context.BindJSON(&newReceipt); err != nil {
		context.Error(err)
		return
	}

	id := uuid.New().String()

	data.UpdateReceiptScore(id, 42)

	context.JSON(http.StatusOK, receiptId{id})
}

// Return the point value for the receipt with the given ID.
func GetPointsForReceipt(context *gin.Context) {
	id := context.Param("id")

	points, present := data.GetScoreForReceiptId(id)

	if !present {
		context.String(http.StatusNotFound, "No record found for receipt %s", id)
		return
	}

	context.JSON(http.StatusOK, receiptPoints{points})
}
