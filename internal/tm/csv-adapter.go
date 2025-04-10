package tm

import (
	"encoding/csv"
	"fmt"
	"io"
)

type CsvFileAdapter func(rows [][]string) ([]importTransactionParams, error)

var adapters = []CsvFileAdapter{
	ingFileAdapter,
	bendigoFileAdapter,
}

func ParseCsvFile(f io.Reader) ([]importTransactionParams, error) {
	r := csv.NewReader(f)

	/*
	   Should stream this in the future, however atm it meets the requirements because
	   the csv files I upload every month are small.
	*/
	rows, err := r.ReadAll()
	if err != nil {
		return nil, UserErr(fmt.Sprintf("Error parsing file as csv: %s", err))
	}

	errs := []error{}
	for _, adapter := range adapters {
		params, err := adapter(rows)
		if err != nil {
			errs = append(errs, err)
			continue
		}
		return params, nil
	}

	/*
	   Should surface better errors to the user, however atm I don't care
	*/
	return nil, UserErr(fmt.Sprintf("Could not find suitable file adapter errors: %s", errs))
}
