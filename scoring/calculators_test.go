package scoring

import (
	"jjcaptain/receipt-processor/types"
	"testing"
)

func TestScoreRetailer_SimpleCase(t *testing.T) {
	input := types.Receipt{Retailer: "Target"}
	expected := 6

	score, _ := ScoreRetailer(input)

	if score != expected {
		t.Errorf("Score %d did not match expected %d", score, expected)
	}
}

func TestScoreRetailer_WithSymbols(t *testing.T) {
	input := types.Receipt{Retailer: "M&M Corner Market"}
	expected := 14

	score, _ := ScoreRetailer(input)

	if score != expected {
		t.Errorf("Score %d did not match expected %d", score, expected)
	}
}

func TestScoreTotal_ValueWithNondivisibleCents(t *testing.T) {
	input := types.Receipt{Total: "35.35"}
	expected := 0

	score, err := ScoreTotal(input)

	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if score != expected {
		t.Errorf("Score %d did not match expected %d", score, expected)
	}
}

func TestScoreTotal_ValueWithDivisibleCents(t *testing.T) {
	input := types.Receipt{Total: "22.25"}
	expected := 25

	score, err := ScoreTotal(input)

	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if score != expected {
		t.Errorf("Score %d did not match expected %d", score, expected)
	}
}

func TestScoreTotal_ValueWithWholeDollar(t *testing.T) {
	input := types.Receipt{Total: "12.00"}
	expected := 75

	score, err := ScoreTotal(input)

	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if score != expected {
		t.Errorf("Score %d did not match expected %d", score, expected)
	}
}

func TestScoreTotal_InvalidValue(t *testing.T) {
	input := types.Receipt{Total: "notAnumber"}

	_, err := ScoreTotal(input)

	if err == nil {
		t.Error("Expected error but did not get one")
	}
}

func TestScoreItems_EvenItemCount(t *testing.T) {
	input := types.Receipt{Items: []types.ReceiptItem{
		{ShortDescription: "Gatorade", Price: "2.25"},
		{ShortDescription: "Gatorade", Price: "2.25"},
		{ShortDescription: "Gatorade", Price: "2.25"},
		{ShortDescription: "Gatorade", Price: "2.25"}} }
	expected := 10

	score, err := ScoreItems(input)

	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if score != expected {
		t.Errorf("Score %d did not match expected %d", score, expected)
	}
}

func TestScoreItems_OddItemCount(t *testing.T) {
	input := types.Receipt{Items: []types.ReceiptItem{
		{ShortDescription: "Gatorade", Price: "2.25"},
		{ShortDescription: "Gatorade", Price: "2.25"},
		{ShortDescription: "Gatorade", Price: "2.25"},
		{ShortDescription: "Gatorade", Price: "2.25"},
		{ShortDescription: "Gatorade", Price: "2.25"}} }
	expected := 10

	score, err := ScoreItems(input)

	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if score != expected {
		t.Errorf("Score %d did not match expected %d", score, expected)
	}
}

func TestScoreItems_ScorableDescription(t *testing.T) {
	input := types.Receipt{Items: []types.ReceiptItem{
		{ShortDescription: "Emils Cheese Pizza", Price: "12.25"}} }
	expected := 3

	score, err := ScoreItems(input)

	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if score != expected {
		t.Errorf("Score %d did not match expected %d", score, expected)
	}
}

func TestScoreItems_DescriptionWithExtraSpaces(t *testing.T) {
	input := types.Receipt{Items: []types.ReceiptItem{
		{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"}} }
	expected := 3

	score, err := ScoreItems(input)

	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if score != expected {
		t.Errorf("Score %d did not match expected %d", score, expected)
	}
}

func TestScoreItems_MeetsMultipleCriteria(t *testing.T) {
	input := types.Receipt{Items: []types.ReceiptItem{
		{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
		{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
		{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
		{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
		{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"}} }

	expected := 16

	score, err := ScoreItems(input)

	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if score != expected {
		t.Errorf("Score %d did not match expected %d", score, expected)
	}
}

func TestScoreItems_PriceNotANumber(t *testing.T) {
	input := types.Receipt{Items: []types.ReceiptItem{
		{ShortDescription: "Emils Cheese Pizza", Price: "pizza"}} }

	_, err := ScoreItems(input)

	if err == nil {
		t.Error("Expected error but did not get one")
	}
}

func TestScoreDate_OddDay(t *testing.T) {
	input := types.Receipt{PurchaseDate: "2023-05-15"}
	expected := 6

	score, err := ScoreDate(input)

	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if score != expected {
		t.Errorf("Score %d did not match expected %d", score, expected)
	}
}

func TestScoreDate_EvenDay(t *testing.T) {
	input := types.Receipt{PurchaseDate: "2023-05-16"}
	expected := 0

	score, err := ScoreDate(input)

	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if score != expected {
		t.Errorf("Score %d did not match expected %d", score, expected)
	}
}

func TestScoreDate_InvalidDay(t *testing.T) {
	input := types.Receipt{PurchaseDate: "notAdate"}

	_, err := ScoreDate(input)

	if err == nil {
		t.Error("Expected error but did not get one")
	}
}

func TestScoreTime_TimeTooEarly(t *testing.T) {
	input := types.Receipt{PurchaseTime: "13:01"}
	expected := 0

	score, err := ScoreTime(input)

	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if score != expected {
		t.Errorf("Score %d did not match expected %d", score, expected)
	}
}

func TestScoreTime_TimeTooLate(t *testing.T) {
	input := types.Receipt{PurchaseTime: "17:01"}
	expected := 0

	score, err := ScoreTime(input)

	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if score != expected {
		t.Errorf("Score %d did not match expected %d", score, expected)
	}
}

func TestScoreTime_TimeInRange(t *testing.T) {
	input := types.Receipt{PurchaseTime: "14:33"}
	expected := 10

	score, err := ScoreTime(input)

	if err != nil {
		t.Error("Unexpected error:", err)
	}
	if score != expected {
		t.Errorf("Score %d did not match expected %d", score, expected)
	}
}

func TestScoreTime_InvalidTime(t *testing.T) {
	input := types.Receipt{PurchaseTime: "notAtime"}

	_, err := ScoreTime(input)

	if err == nil {
		t.Error("Expected error but did not get one")
	}
}
