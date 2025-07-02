package tm_test

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"
	"tm/internal/cfg"
	"tm/internal/orm"
	"tm/internal/orm/model"
	"tm/internal/tm"

	"github.com/stretchr/testify/require"
)

type csvFileCase struct {
	name     string
	filePath string
	expected map[string]model.TmTransaction
}

var csvFileFormats = []csvFileCase{
	{
		name:     "ing",
		filePath: "../../fixtures/ing_transactions.csv",
		expected: map[string]model.TmTransaction{
			"12/12/2023Credit": {
				ID:          "12/12/2023Credit",
				Date:        time.Date(2023, 12, 12, 0, 0, 0, 0, time.UTC),
				Description: "Credit",
				AmountCents: 3000,
			},
			"24/01/2024Some buy hey?": {
				ID:          "24/01/2024Some buy hey?",
				Date:        time.Date(2024, 1, 24, 0, 0, 0, 0, time.UTC),
				Description: "Some buy hey?",
				AmountCents: -4800,
			},
		},
	},
	{
		name:     "bendigo",
		filePath: "../../fixtures/bendigo_transactions.csv",
		expected: map[string]model.TmTransaction{
			"31/01/2024Description 1": {
				ID:          "31/01/2024Description 1",
				Date:        time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC),
				Description: "Description 1",
				AmountCents: 50000,
			},
			"31/01/2024Description 2": {
				ID:          "31/01/2024Description 2",
				Date:        time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC),
				Description: "Description 2",
				AmountCents: -1310,
			},
		},
	},
}

func TestImportCsv(t *testing.T) {
	gormDb, err := orm.NewGormDb(cfg.Must())
	if err != nil {
		require.NoError(t, err)
	}
	manager := tm.New(gormDb)
	ctx := context.Background()

	for _, csvFileFormat := range csvFileFormats {
		t.Run(fmt.Sprintf("Should save %s transactions to the database", csvFileFormat.name), func(t *testing.T) {
			t.Cleanup(func() {
				gormDb.Where("1=1").Delete(&model.TmTransaction{})
			})

			f, err := os.Open(csvFileFormat.filePath)
			require.NoError(t, err)
			t.Cleanup(func() {
				f.Close()
			})

			res, err := manager.ImportCsv(ctx, f)
			require.NoError(t, err)
			require.Equal(t, len(csvFileFormat.expected), res.Total)

			var transactions []model.TmTransaction
			err = gormDb.Find(&transactions).Error
			require.NoError(t, err)
			require.Len(t, transactions, 2)
			for _, transaction := range transactions {
				require.Equal(t, csvFileFormat.expected[transaction.ID], transaction)
			}
		})
	}

	t.Run("Should not override existing duplicate transactions", func(t *testing.T) {
		t.Cleanup(func() {
			gormDb.Where("1=1").Delete(&model.TmTransaction{})
		})

		existingRow := model.TmTransaction{
			ID:          "31/01/2024Description 1",
			Date:        time.Date(2024, 1, 31, 0, 0, 0, 0, time.UTC),
			Description: "original description",
			AmountCents: 40000,
		}

		err := gormDb.Create(&existingRow).Error
		require.NoError(t, err)

		testCsvFile := `"31/01/2024","500.00","Description 1"`
		reader := strings.NewReader(testCsvFile)

		res, err := manager.ImportCsv(context.Background(), reader)
		require.NoError(t, err)
		require.Equal(t, 1, res.Total)
		require.Equal(t, 1, res.Duplicates)

		var resultTransaction model.TmTransaction
		err = gormDb.Where("id = ?", existingRow.ID).First(&resultTransaction).Error
		require.NoError(t, err)
		require.Equal(t, existingRow, resultTransaction)
	})
}
