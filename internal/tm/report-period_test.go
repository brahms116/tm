package tm

import (
	"context"
	"testing"
	"time"
	"tm/internal/cfg"
	"tm/internal/orm"
	"tm/internal/orm/model"
	"tm/pkg/contracts"

	"github.com/stretchr/testify/require"
)

func requireContractEqualModel(t *testing.T, expected model.TmTransaction, got contracts.Transaction) {
	t.Helper()
	require.Equal(t, expected.ID, got.Id)
	require.Equal(t, expected.Date.Format(time.RFC3339), got.Date)
	require.Equal(t, expected.Description, got.Description)
	require.Equal(t, int(expected.AmountCents), got.AmountCents)
	require.Equal(t, expected.CategoryID, got.Category)
}

func TestReportPeriod(t *testing.T) {
	gormDb, err := orm.NewGormDb(cfg.Must())
	if err != nil {
		require.NoError(t, err)
	}
	tm := New(gormDb)
	ctx := context.Background()

	t.Cleanup(func() {
		gormDb.Where("1=1").Delete(&model.TmTransaction{})
	})

	start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, 12, 31, 0, 0, 0, 0, time.UTC)

	models := []model.TmTransaction{
		{
			ID:          "before",
			Date:        start.AddDate(0, 0, -1),
			AmountCents: 1000,
			Description: "Before",
		},
		{
			ID:          "during1",
			Date:        start,
			AmountCents: 1000,
			Description: "During 1",
		},
		{
			ID:          "during2",
			Date:        start.AddDate(0, 0, 1),
			AmountCents: 1800,
			Description: "During 2",
		},
		{
			ID:          "during3",
			Date:        start.AddDate(0, 0, 2),
			AmountCents: -900,
			Description: "During 3",
		},
		{
			ID:          "after",
			Date:        end,
			AmountCents: 1000,
			Description: "after",
		},
	}

	idToModel := map[string]model.TmTransaction{}
	for _, model := range models {
		idToModel[model.ID] = model
	}

	gormDb.Create(&models)
	result, err := tm.ReportPeriod(ctx, start, end, false)
	require.NoError(t, err)

	t.Run("Should calculate the correct stats for the period", func(t *testing.T) {
		require.Equal(t, 2800, result.Summary.EarningCents)
		require.Equal(t, 900, result.Summary.SpendingCents)
		require.Equal(t, 1900, result.Summary.NetCents)
	})

	t.Run("Should return the correct spending transactions", func(t *testing.T) {
		require.Len(t, result.TopSpendings, 1)
		requireContractEqualModel(t, idToModel["during3"], result.TopSpendings[0])
	})

	t.Run("Should return the correct earning transactions", func(t *testing.T) {
		require.Len(t, result.TopEarnings, 2)

		idToResult := map[string]contracts.Transaction{}
		for _, transaction := range result.TopEarnings {
			idToResult[transaction.Id] = transaction
		}

		earning1, ok := idToResult["during1"]
		require.True(t, ok)

		earning2, ok := idToResult["during2"]
		require.True(t, ok)

		requireContractEqualModel(t, idToModel["during1"], earning1)
		requireContractEqualModel(t, idToModel["during2"], earning2)
	})
}
