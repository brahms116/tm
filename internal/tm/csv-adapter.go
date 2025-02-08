package tm

import (
	"encoding/csv"
	"fmt"
	"io"
	"tm/internal/data"
)

type CsvFileAdapter interface {
	Parse(rows [][]string) ([]data.AddTransactionParams, error)
}

var adapters = []CsvFileAdapter{
	IngFileAdapter{},
	BendigoCsvRowAdapter{},
}

func ParseCsvFile(f io.Reader) ([]data.AddTransactionParams, error) {
	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading csv: %w", err)
	}
	errors := []error{}

	for _, adapter := range adapters {
		params, err := adapter.Parse(rows)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		return params, nil
	}
	return nil, fmt.Errorf("error parsing csv: %v", errors)
}
