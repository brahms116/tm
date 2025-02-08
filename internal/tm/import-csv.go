package tm

import (
	"context"
	"fmt"
	"io"
	"tm/internal/data"
	"tm/pkg/contracts"
)

func (t *tm) ImportCsv(ctx context.Context, f io.Reader) (contracts.ImportCsvResponse, error) {
	params, err := ParseCsvFile(f)
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
		Total:      len(params),
		Duplicates: duplicatesCount,
	}, nil
}
