package main

import (
	"fmt"
	"io"
	"net/http"
	"tm/internal/db"
	"tm/internal/tm"
)

func main() {
	db, err := db.New()
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	mux.HandleFunc("POST /import/ing", func(w http.ResponseWriter, r *http.Request) {
		f, close, ok := ReadFile(w, r, "file")
		if !ok {
			return
		}
		defer close()

		dups, err := tm.LoadIngCsv(db, f)
		if err != nil {
			panic(err)
		}
		w.Write([]byte(fmt.Sprintf("Dups: %d", dups)))
	})

	err = http.ListenAndServe(":8081", mux)
	if err != nil {
		panic(err)
	}
}

func ReadFile(w http.ResponseWriter, r *http.Request, key string) (io.Reader, func() error, bool) {
	f, _, err := r.FormFile(key)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return nil, nil, false
	}
	return f, f.Close, true
}
