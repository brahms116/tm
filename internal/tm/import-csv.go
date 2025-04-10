package tm

import (
	"context"
	"gorm.io/gorm/clause"
	"io"
	"tm/pkg/contracts"
)

func (t *tm) ImportCsv(ctx context.Context, f io.Reader) (contracts.ImportCsvResponse, error) {
	params, err := ParseCsvFile(f)
	if err != nil {
		return contracts.ImportCsvResponse{}, wrapErr(err, "error parsing csv file")
	}

	duplicatesCount := 0

	for _, param := range params {
		dbModel := param.toDbModel()
		result := t.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbModel)
		if result.Error != nil {
			return contracts.ImportCsvResponse{}, wrapErr(result.Error, "error inserting transaction")
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
