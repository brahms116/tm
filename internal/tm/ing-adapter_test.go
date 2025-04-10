package tm

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestIngFileAdapter(t *testing.T) {
	expected := []importTransactionParams{
		{
			id:          "24/01/2024Some buy hey?",
			date:        time.Date(2024, 1, 24, 0, 0, 0, 0, time.UTC),
			description: "Some buy hey?",
			amountCents: -4800,
		},
		{
			id:          "12/12/2023Credit",
			date:        time.Date(2023, 12, 12, 0, 0, 0, 0, time.UTC),
			description: "Credit",
			amountCents: 3000,
		},
	}

	f, err := os.Open("../../fixtures/ing_transactions.csv")
	if err != nil {
		t.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	got, err := ParseCsvFile(f)
	if err != nil {
		t.Fatalf("error parsing csv: %v", err)
	}

	assert.Equal(t, expected, got)
}
