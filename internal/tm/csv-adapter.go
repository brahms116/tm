package tm

import (
	"encoding/csv"
	"fmt"
	"io"
	"tm/internal/data"
)

type CsvFileAdapter func(file io.Reader) ([]data.AddTransactionParams, error)

type CsvRowAdapter func(row []string) (data.AddTransactionParams, error)

func NewCsvFileAdapter(adapter CsvRowAdapter) CsvFileAdapter {
	return func(f io.Reader) ([]data.AddTransactionParams, error) {
		r := csv.NewReader(f)
		addParams := []data.AddTransactionParams{}
		rows, err := r.ReadAll()
		if err != nil {
			return nil, fmt.Errorf("error reading csv: %w", err)
		}
		rows = rows[1:] // skip header
		for _, row := range rows {
			addParam, err := adapter(row)
			if err != nil {
				return nil, fmt.Errorf("error parsing row: %v : %w", row, err)
			}
			addParams = append(addParams, addParam)
		}
		return addParams, nil
	}
}

