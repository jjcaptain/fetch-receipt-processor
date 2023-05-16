package data

import "testing"

func TestBasicSaveAndGet(t *testing.T) {
	id := "test-id"
	score := 42

	UpdateReceiptScore(id, score)

	result, present := GetScoreForReceiptId(id)

	if !present {
		t.Error("Record was not present")
	}
	if result != score {
		t.Errorf("Score %d did not match expected %d", result, score)
	}
}

func TestMultiSave(t *testing.T) {
	id := "test-multi-save-id"
	score1 := 42
	score2 := 9000

	UpdateReceiptScore(id, score1)
	UpdateReceiptScore(id, score2)

	result, present := GetScoreForReceiptId(id)

	if !present {
		t.Error("Record was not present")
	}
	if result != score2 {
		t.Errorf("Score %d did not match expected %d", result, score2)
	}
}

func TestXxx(t *testing.T) {
	id := "test-no-record-id"

	_, present := GetScoreForReceiptId(id)

	if present {
		t.Error("Unexpected record was present")
	}
}
