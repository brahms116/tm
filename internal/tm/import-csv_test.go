package tm

import (
	"fmt"
	"testing"
	"time"
	"tm/internal/cfg"
	"tm/internal/orm"
	"tm/internal/orm/model"

	"github.com/stretchr/testify/require"
)

func TestImportCsv(t *testing.T) {
	gormDb, err := orm.NewGormDb(cfg.Must())
	if err != nil {
		require.NoError(t, err)
	}
	t.Cleanup(func() {
		gormDb.Where("1=1").Delete(&model.TmTransaction{})
	})
	// transactionManager := New(nil, gormDb)

	transaction := model.TmTransaction{
		ID:          "31/01/2024Description 1",
		Date:        time.Now(),
		Description: "Description 1",
		AmountCents: 50000,
	}

	result := gormDb.Create(&transaction)
	require.NoError(t, result.Error)

	var transactions []model.TmTransaction
	result = gormDb.Find(&transactions)
	require.NoError(t, result.Error)

	fmt.Printf("Transactions: %v\n", transactions)
}
