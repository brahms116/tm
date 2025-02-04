package tm

import (
	"context"
	"io"
	"tm/internal/db"
)

type TM interface {
	ImportIngCsv(ctx context.Context, f io.Reader) (int, error)
}

type tm struct {
	conn *db.DB
}

func New(conn *db.DB) TM {
	return &tm{conn: conn}
}

