package tm

import (
	"context"
	"io"
	"time"
	"tm/internal/db"
	"tm/pkg/contracts"
)

type TM interface {
	ImportCsv(ctx context.Context, f io.Reader) (contracts.ImportCsvResponse, error)
  Report(ctx context.Context, dateWeek time.Time, numWeeks int) (contracts.ReportResponse, error)
}

type tm struct {
	conn *db.DB
}

func New(conn *db.DB) TM {
	return &tm{conn: conn}
}
