package tm

import (
	"context"
	"gorm.io/gorm"
	"io"
	"time"
	"tm/internal/db"
	"tm/pkg/contracts"
)

type TM interface {
	ImportCsv(ctx context.Context, f io.Reader) (contracts.ImportCsvResponse, error)
	ReportTimeline(ctx context.Context, start, end time.Time) (contracts.TimelineResponse, error)
	ReportPeriod(ctx context.Context, start, end time.Time, u100 bool) (contracts.ReportResponse, error)
}

type tm struct {
	conn *db.DB
	db   *gorm.DB
}

func New(conn *db.DB, db *gorm.DB) TM {
	return &tm{conn: conn, db: db}
}
