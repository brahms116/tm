package tm

import (
	"testing"
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
	tm := New(gormDb)
}
