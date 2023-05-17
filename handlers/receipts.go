package handlers

import (
	"jjcaptain/receipt-processor/data"
	"jjcaptain/receipt-processor/scoring"
	"jjcaptain/receipt-processor/types"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var scorers = []func(types.Receipt) (int, error){scoring.ScoreRetailer, scoring.ScoreTotal, scoring.ScoreItems, scoring.ScoreDate, scoring.ScoreTime}

type receiptId struct {
	ID string `json:"id"`
}

type receiptPoints struct {
	Points int `json:"points"`
}

// Process a single receipt. Generate an ID, calculate and store the amount of points the receipt
// is worth, and return the generated ID.
func ProcessReceipt(context *gin.Context) {
	var receipt types.Receipt

	if err := context.BindJSON(&receipt); err != nil {
		context.Error(err)
		return
	}

	score, err := calculateScore(scorers, receipt)
	if err != nil {
		context.Error(err)
		return
	}

	id := uuid.New().String()
	data.UpdateReceiptScore(id, score)

	context.JSON(http.StatusOK, receiptId{id})
}

func calculateScore(scorers []func(types.Receipt) (int, error), receipt types.Receipt) (int, error) {
	numScorers := len(scorers)

	ch := make(chan int, numScorers)
	errs := make(chan error, 1)

	// spin up all the scorers
	for i := range scorers {
		go func(scorer func(types.Receipt) (int, error)) {
			subScore, err := scorer(receipt)
			if err != nil {
				errs <- err
			} else {
				ch <- subScore
			}
		}(scorers[i])
	}

	score := 0
	// add up the results
	var subScore int
	var err error
	for i := 0; i < numScorers; i++ {
		select {
		case subScore = <- ch:
			score += subScore
		case err = <- errs:
			return 0, err
		}
	}
	
	return score, nil
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
