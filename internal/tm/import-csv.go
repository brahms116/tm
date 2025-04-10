package tm

import (
	"context"
	"fmt"
	"io"
	"tm/pkg/contracts"
	"gorm.io/gorm/clause"
)

func (t *tm) ImportCsv(ctx context.Context, f io.Reader) (contracts.ImportCsvResponse, error) {
	params, err := ParseCsvFile(f)
	if err != nil {
		return contracts.ImportCsvResponse{}, fmt.Errorf("error parsing csv: %w", err)
	}

	duplicatesCount := 0

	for _, param := range params {
		dbModel := param.toDbModel()
		result := t.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbModel)
		if result.Error != nil {
			return contracts.ImportCsvResponse{}, fmt.Errorf("error adding transaction: %w", err)
		}
		if result.RowsAffected != 1 {
			duplicatesCount++
		}
	}

	return contracts.ImportCsvResponse{
		Total:      len(params),
		Duplicates: duplicatesCount,
	}, nil
}
