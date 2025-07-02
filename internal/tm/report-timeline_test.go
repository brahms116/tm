package tm_test

import (
	"context"
	"testing"
	"time"
	"tm/internal/cfg"
	"tm/internal/orm"
	"tm/internal/orm/model"
	"tm/internal/tm"

	"github.com/stretchr/testify/require"
)

func TestReportTimeline(t *testing.T) {
	gormDb, err := orm.NewGormDb(cfg.Must())
	if err != nil {
		require.NoError(t, err)
	}
	manager := tm.New(gormDb)
	ctx := context.Background()

	t.Cleanup(func() {
		gormDb.Where("1=1").Delete(&model.TmTransaction{})
	})

	start := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, 3, 1, 0, 0, 0, 0, time.UTC)

	models := []model.TmTransaction{
		{
			ID:          "jan1",
			Date:        start.AddDate(0, 0, 1),
			AmountCents: 1000,
			Description: "jan1",
		},
		{
			ID:          "jan2",
			Date:        start.AddDate(0, 0, 2),
			AmountCents: -400,
			Description: "jan2",
		},
		{
			ID:          "jan3",
			Date:        start.AddDate(0, 0, 3),
			AmountCents: 200,
			Description: "jan3",
		},
		{
			ID:          "feb1",
			Date:        start.AddDate(0, 1, 2),
			AmountCents: -900,
			Description: "feb1",
		},
		{
			ID:          "feb2",
			Date:        start.AddDate(0, 1, 3),
			AmountCents: -800,
			Description: "feb2",
		},
		{
			ID:          "feb3",
			Date:        start.AddDate(0, 1, 0),
			AmountCents: 1000,
			Description: "feb3",
		},
		{
			ID:          "mar1",
			Date:        start.AddDate(0, 2, 0),
			AmountCents: 1000,
			Description: "mar1",
		},
	}

	idToModel := map[string]model.TmTransaction{}
	for _, model := range models {
		idToModel[model.ID] = model
	}
	gormDb.Create(&models)
	result, err := manager.ReportTimeline(ctx, start, end)
	require.NoError(t, err)

	require.Len(t, result.Items, 2)

  require.Equal(t, start.Format(time.RFC3339), result.Items[0].Month)
  require.Equal(t, 400, result.Items[0].Summary.SpendingCents)
  require.Equal(t, 1200, result.Items[0].Summary.EarningCents)
  require.Equal(t, 800, result.Items[0].Summary.NetCents)


  require.Equal(t, start.AddDate(0, 1, 0).Format(time.RFC3339), result.Items[1].Month)
  require.Equal(t, 1700, result.Items[1].Summary.SpendingCents)
  require.Equal(t, 1000, result.Items[1].Summary.EarningCents)
  require.Equal(t, -700, result.Items[1].Summary.NetCents)
}
