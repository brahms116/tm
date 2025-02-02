package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*pgxpool.Pool
}

func New() (*DB, error) {
	conn, err := pgxpool.New(context.TODO(), "postgres://postgres:password@localhost:5432/tm_dev")
	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}
	return &DB{conn}, nil
}
