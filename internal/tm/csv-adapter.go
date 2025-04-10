package tm

import (
	"encoding/csv"
	"fmt"
	"io"
)

type CsvFileAdapter func (rows [][]string) ([]importTransactionParams, error)

var adapters = []CsvFileAdapter{
	ingFileAdapter,
	bendigoFileAdapter,
}

func ParseCsvFile(f io.Reader) ([]importTransactionParams, error) {
	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading csv: %w", err)
	}
	errors := []error{}

	for _, adapter := range adapters {
		params, err := adapter(rows)
		if err != nil {
			errors = append(errors, err)
			continue
		}
		return params, nil
	}
	return nil, fmt.Errorf("error parsing csv: %v", errors)
}
