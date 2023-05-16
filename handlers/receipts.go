package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type receiptId struct {
	ID string `json:"id"`
}

type receiptPoints struct {
	Points int `json:"points"`
}

func ProcessReceipt(context *gin.Context) {

	context.JSON(http.StatusOK, receiptId{"example-id"})
}

func GetPointsForReceipt(context *gin.Context) {

	context.JSON(http.StatusOK, receiptPoints{42})
}