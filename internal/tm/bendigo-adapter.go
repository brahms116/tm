package tm

import (
	"fmt"
	"strconv"
	"time"
)

func bendigoFileAdapter(rows [][]string) ([]importTransactionParams, error) {
	addParams := []importTransactionParams{}
	for _, row := range rows {
		addParam, err := bendigoRowAdapter(row)
		if err != nil {
			return nil, fmt.Errorf("error parsing Bendigo row: %v : %w", row, err)
		}
		addParams = append(addParams, addParam)
	}
	return addParams, nil
}

// Parses a row like "31/01/2024","500.00","Description 1"
func bendigoRowAdapter(row []string) (importTransactionParams, error) {
	if len(row) != 3 {
		return importTransactionParams{}, fmt.Errorf("expected 3 columns, got %d", len(row))
	}

	dateStr := row[0]
	amountStr := row[1]
	description := row[2]

	id := dateStr + description

	time, err := time.Parse("02/01/2006", dateStr)
	if err != nil {
		return importTransactionParams{}, fmt.Errorf("error parsing date: %w", err)
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return importTransactionParams{}, fmt.Errorf("error parsing amount: %w", err)
	}
	amountCents := int32(amount * 100)

	return importTransactionParams{
		id:          id,
		date:        time,
		description: description,
		amountCents: amountCents,
	}, nil
}
