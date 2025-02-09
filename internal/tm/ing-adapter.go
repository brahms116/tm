package tm

import (
	"fmt"
	"strconv"
	"time"
	"tm/internal/data"
)

func ingFileAdapter(rows [][]string) ([]data.AddTransactionParams, error) {
	header := rows[0]
	if len(header) != 5 {
		return nil, fmt.Errorf("expected 5 columns, got %d", len(header))
	}
	if header[0] != "Date" ||
		header[1] != "Description" ||
		header[2] != "Credit" ||
		header[3] != "Debit" ||
		header[4] != "Balance" {
		return nil, fmt.Errorf("unexpected header: %v", header)
	}

	rows = rows[1:] // skip header
	addParams := []data.AddTransactionParams{}
	for _, row := range rows {
		addParam, err := ingRowAdapter(row)
		if err != nil {
			return nil, fmt.Errorf("error parsing Ing row: %v : %w", row, err)
		}
		addParams = append(addParams, addParam)
	}
	return addParams, nil
}

func ingRowAdapter(row []string) (data.AddTransactionParams, error) {
	dateStr := row[0]
	description := row[1]
	creditStr := row[2]
	debitStr := row[3]

	var amountStr string
	if creditStr != "" {
		amountStr = creditStr
	} else {
		amountStr = debitStr
	}

	id := dateStr + description

	time, err := time.Parse("02/01/2006", dateStr)
	if err != nil {
		return data.AddTransactionParams{}, fmt.Errorf("error parsing date: %w", err)
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return data.AddTransactionParams{}, fmt.Errorf("error parsing amount: %w", err)
	}
	amountCents := int32(amount * 100)

	return data.AddTransactionParams{
		ID:          id,
		Date:        time,
		Description: description,
		AmountCents: amountCents,
	}, nil
}
