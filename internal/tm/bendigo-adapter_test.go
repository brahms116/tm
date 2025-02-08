package tm

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
	"tm/internal/data"
)

func TestBendigoFileAdapter(t *testing.T) {
	expected := []data.AddTransactionParams{
		{
			ID:          "31/01/2024Description 1",
			Date:        time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC),
			Description: "Description 1",
			AmountCents: 50000,
		},
		{
			ID:          "31/01/2024Description 2",
			Date:        time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC),
			Description: "Description 2",
			AmountCents: -1310,
		},
	}

	f, err := os.Open("../../fixtures/bendigo_transactions.csv")
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
