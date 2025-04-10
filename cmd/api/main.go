package main

import (
	"tm/internal/cfg"
	"tm/internal/db"
	"tm/internal/orm"
	"tm/internal/tm"
)

func main() {
	cfg := cfg.Must()
	db, err := db.New(cfg)
	if err != nil {
		panic(err)
	}

	gormDb, err := orm.NewGormDb(cfg)
	if err != nil {
		panic(err)
	}

	transactionManager := tm.New(db, gormDb)
	server := New(cfg, transactionManager)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
