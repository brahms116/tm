package tm

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"time"
	"tm/internal/data"
	"tm/internal/db"
)

func LoadIngCsv(f io.Reader, conn *db.DB) (int, error) {
	params, err := ParseIngCsv(f)
	if err != nil {
		return 0, fmt.Errorf("error parsing csv: %w", err)
	}

	duplicatesCount := 0

	for _, param := range params {
		count, err := data.New().AddTransaction(context.TODO(), conn, param)
		if err != nil {
			return 0, fmt.Errorf("error adding transaction: %w", err)
		}
		if count != 1 {
			duplicatesCount++
		}
	}

	return duplicatesCount, nil
}

func ParseIngCsv(f io.Reader) ([]data.AddTransactionParams, error) {
	r := csv.NewReader(f)
	addParams := []data.AddTransactionParams{}
	rows, err := r.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading csv: %w", err)
	}
	rows = rows[1:] // skip header
	for _, row := range rows {
		addParam, err := parseIngRow(row)
		if err != nil {
			return nil, fmt.Errorf("error parsing row: %v : %w", row, err)
		}
		addParams = append(addParams, addParam)
	}
	return addParams, nil
}

func parseIngRow(row []string) (data.AddTransactionParams, error) {
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
