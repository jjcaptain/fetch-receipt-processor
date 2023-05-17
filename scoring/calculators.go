package scoring

import (
	"jjcaptain/receipt-processor/types"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Calculate points for the retailer
func ScoreRetailer(receipt types.Receipt) (int, error) {
	// Award one point per alphanumeric character
	nonAlphanumericRegex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	filteredRetailerName := nonAlphanumericRegex.ReplaceAllString(receipt.Retailer, "")
	return len(filteredRetailerName), nil
}

// Calculate points for the receipt total
func ScoreTotal(receipt types.Receipt) (int, error) {
	// First make sure it's a number
	total, err := strconv.ParseFloat(receipt.Total, 32)
	if err != nil {
		return 0, err
	}

	// Multiply by 100 and convert to int to make the comparisons cleaner
	boostedTotal := int(total * 100)

	score := 0
	// 50 points if the total is a round dollar amount with no cents
	if boostedTotal % 100 == 0 {
		score += 50
	}
	// 25 points if the total is a multiple of 0.25
	if boostedTotal % 25 == 0 {
		score += 25
	}
	
	return score, nil
}

// Calculate points for items on the receipt
func ScoreItems(receipt types.Receipt) (int, error) {
	score := 0

	numItems := len(receipt.Items)
	// 5 points for every two items on the receipt
	score += (5 * (numItems / 2))

	ch := make(chan int, numItems)
	errs := make(chan error, 1)

	// calculate the score for each item
	for i := range receipt.Items {
		go func (item types.ReceiptItem)  {
			subScore, err := scoreItem(item)
			if err != nil {
				errs <- err
			} else {
				ch <- subScore
			}
		}(receipt.Items[i])
	}

	// add up the results
	var itemScore int
	var err error
	for i := 0; i < numItems; i++ {
		select {
		case itemScore = <- ch:
			score += itemScore
		case err = <- errs:
			return 0, err
		}
	}

	return score, nil
}

// Calculate points for a single item on the receipt
func scoreItem(item types.ReceiptItem) (int, error) {
	score := 0

	// If the trimmed length of the item description is a multiple of 3
	trimmedDescription := strings.TrimSpace(item.ShortDescription)
	if len(trimmedDescription) % 3 == 0 {
		price, err := strconv.ParseFloat(item.Price, 32)
		if err != nil {
			return 0, err
		}

		// multiply the price by 0.2 and round up to the nearest integer
		score = int(math.Ceil(price * 0.2))
	}

	return score, nil
}

// Calculate points for the receipt date
func ScoreDate(receipt types.Receipt) (int, error) {
	purchaseDate, err := time.Parse(time.DateOnly, receipt.PurchaseDate)
	if err != nil {
		return 0, err
	}

	score := 0
	// 6 points if the day in the purchase date is odd
	if purchaseDate.Day() % 2 == 1 {
		score = 6
	}

	return score, nil
}

// Calculate points for the receipt time
func ScoreTime(receipt types.Receipt) (int, error) {
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err != nil {
		return 0, err
	}

	score := 0
	// 10 points if the time of purchase is after 2:00pm and before 4:00pm
	startTime := time.Date(purchaseTime.Year(), purchaseTime.Month(), purchaseTime.Day(), 14, 0, 0, 0, purchaseTime.Location())
	endTime := time.Date(purchaseTime.Year(), purchaseTime.Month(), purchaseTime.Day(), 16, 0, 0, 0, purchaseTime.Location())

	if purchaseTime.After(startTime) && purchaseTime.Before(endTime) {
		score = 10
	}

	return score, nil
}
