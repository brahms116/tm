package tm

import (
	"os"
	"testing"
	"time"
	"tm/internal/data"
	"github.com/stretchr/testify/assert"
)

func TestIngFileAdapter(t *testing.T) {
	expected := []data.AddTransactionParams{
		{
			ID:          "24/01/2024Some buy hey?",
			Date:        time.Date(2024, 1, 24, 0, 0, 0, 0, time.UTC),
			Description: "Some buy hey?",
			AmountCents: -4800,
		},
		{
			ID:          "12/12/2023Credit",
			Date:        time.Date(2023, 12, 12, 0, 0, 0, 0, time.UTC),
			Description: "Credit",
			AmountCents: 3000,
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
