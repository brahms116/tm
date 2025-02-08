package main

import (
	"tm/internal/cfg"
	"tm/internal/db"
	"tm/internal/tm"
)

func main() {
	cfg := cfg.Must()
	db, err := db.New(cfg)
	if err != nil {
		panic(err)
	}

	transactionManager := tm.New(db)
	server := New(cfg, transactionManager)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
