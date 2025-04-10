package tm

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestBendigoFileAdapter(t *testing.T) {
	expected := []importTransactionParams{
		{
			id:          "31/01/2024Description 1",
			date:        time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC),
			description: "Description 1",
			amountCents: 50000,
		},
		{
			id:          "31/01/2024Description 2",
			date:        time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC),
			description: "Description 2",
			amountCents: -1310,
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
