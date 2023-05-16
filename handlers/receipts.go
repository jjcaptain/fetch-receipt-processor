package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ReceiptItem struct {
	ShortDescription string `json:"shortDescription"`
	Price string `json:"price"`
}

type Receipt struct {
	Retailer string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items []ReceiptItem `json:"items"`
	Total string `json:"total"`
}

type receiptId struct {
	ID string `json:"id"`
}

type receiptPoints struct {
	Points int `json:"points"`
}

func ProcessReceipt(context *gin.Context) {
	var newReceipt Receipt

	if err := context.BindJSON(&newReceipt); err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, receiptId{strconv.Itoa(len(newReceipt.Items))})
}

func GetPointsForReceipt(context *gin.Context) {

	context.JSON(http.StatusOK, receiptPoints{42})
}