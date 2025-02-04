package tm

import (
	"context"
	"fmt"
	"io"
	"tm/internal/data"
	"tm/pkg/contracts"
)

func (t *tm) ImportIngCsv(ctx context.Context, f io.Reader) (contracts.ImportCsvResponse, error) {
	return t.importCsv(ctx, f, NewCsvFileAdapter(IngCsvRowAdapter))
}

// Imports transactions from a CSV file into the database. Returns the number of duplicates.
func (t *tm) importCsv(ctx context.Context, f io.Reader, fileAdapter CsvFileAdapter) (contracts.ImportCsvResponse, error) {
	params, err := fileAdapter(f)
	if err != nil {
		return contracts.ImportCsvResponse{}, fmt.Errorf("error parsing csv: %w", err)
	}

	duplicatesCount := 0

	for _, param := range params {
		count, err := data.New().AddTransaction(ctx, t.conn, param)
		if err != nil {
			return contracts.ImportCsvResponse{}, fmt.Errorf("error adding transaction: %w", err)
		}
		if count != 1 {
			duplicatesCount++
		}
	}

	return contracts.ImportCsvResponse{
		Duplicates: duplicatesCount,
	}, nil
}
