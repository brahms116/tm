package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	*pgxpool.Pool
}

func New() (*DB, error) {
	dbUrl := os.Getenv("TM_DB_URL")
	if dbUrl == "" {
		return nil, fmt.Errorf("TM_DB_URL must be set")
	}
	conn, err := pgxpool.New(context.TODO(), dbUrl)
	if err != nil {
		return nil, fmt.Errorf("error connecting to db: %w", err)
	}
	return &DB{conn}, nil
}
