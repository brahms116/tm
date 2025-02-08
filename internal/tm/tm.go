package tm

import (
	"context"
	"io"
	"tm/internal/db"
	"tm/pkg/contracts"
)

type TM interface {
	ImportCsv(ctx context.Context, f io.Reader) (contracts.ImportCsvResponse, error)
}

type tm struct {
	conn *db.DB
}

func New(conn *db.DB) TM {
	return &tm{conn: conn}
}
