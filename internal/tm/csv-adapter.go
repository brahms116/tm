package tm

import (
	"encoding/csv"
	"fmt"
	"io"
	"tm/internal/data"
)

type CsvFileAdapter interface {
	IsHeader(row []string) bool
	ParseRow(row []string) (data.AddTransactionParams, error)
}

var adapters = []CsvFileAdapter{
	IngFileAdapter{},
}

func pickAdapter(header []string) CsvFileAdapter {
	for _, adapter := range adapters {
		if adapter.IsHeader(header) {
			return adapter
		}
	}
	return nil
}

func ParseCsvFile(f io.Reader) ([]data.AddTransactionParams, error) {
	r := csv.NewReader(f)
	addParams := []data.AddTransactionParams{}
	rows, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading csv: %w", err)
	}
	header := rows[0]
	adapter := pickAdapter(header)
	rows = rows[1:] // skip header
	for _, row := range rows {
		addParam, err := adapter.ParseRow(row)
		if err != nil {
			return nil, fmt.Errorf("error parsing row: %v : %w", row, err)
		}
		addParams = append(addParams, addParam)
	}
	return addParams, nil
}
