package data

var receiptScores = make(map[string]int)

// Add or update the score for a receipt ID in the data store
func UpdateReceiptScore(id string, score int) {
	receiptScores[id] = score
}

// Return score for provided ID if present and indicator of if the score is
// present in the data store
func GetScoreForReceiptId(id string) (int, bool) {
	score, present := receiptScores[id]
	return score, present
}
