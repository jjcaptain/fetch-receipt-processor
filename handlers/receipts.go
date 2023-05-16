package handlers

import (
	"jjcaptain/receipt-processor/types"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

	context.JSON(http.StatusOK, receiptId{strconv.Itoa(len(newReceipt.Items))})
}

// Return the point value for the receipt with the given ID.
func GetPointsForReceipt(context *gin.Context) {

	context.JSON(http.StatusOK, receiptPoints{42})
}
