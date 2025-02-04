package main

import (
	"tm/internal/db"
	"tm/internal/tm"
)

func main() {
	db, err := db.New()
	if err != nil {
		panic(err)
	}

	transactionManager := tm.New(db)
	server := New(transactionManager)
	if err := server.Start(); err != nil {
		panic(err)
	}
}
