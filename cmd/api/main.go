package main

import (
	"tm/internal/cfg"
	"tm/internal/orm"
	"tm/internal/tm"
)

func main() {
	cfg := cfg.Must()
	gormDb, err := orm.NewGormDb(cfg)
	if err != nil {
		panic(err)
	}

	transactionManager := tm.New(gormDb)
	server := New(cfg, transactionManager)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
