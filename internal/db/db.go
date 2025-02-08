package db

import (
	"context"
	"fmt"
	"tm/internal/cfg"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*pgxpool.Pool
}

func New(c cfg.Cfg) (*DB, error) {
	conn, err := pgxpool.New(context.TODO(), c.DbUrl)
	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}
	return &DB{conn}, nil
}
