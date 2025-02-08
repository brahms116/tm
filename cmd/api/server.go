package main

import (
	"fmt"
	"net/http"
	"tm/internal/tm"
	"tm/pkg/handlerutil"
)

type Server struct {
	tm tm.TM
}

func New(tm tm.TM) *Server {
	return &Server{tm: tm}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", s.health)
	mux.HandleFunc("POST /import", s.importIngCsv)

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		return fmt.Errorf("error starting server: %w", err)
	}
	return nil
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	handlerutil.Ok(w, "OK")
}

func (s *Server) importIngCsv(w http.ResponseWriter, r *http.Request) {
	f, close, ok := handlerutil.ReadFile(w, r, "file")
	if !ok {
		return
	}
	defer close()

	result, err := s.tm.ImportCsv(r.Context(), f)
	if err != nil {
		handlerutil.BadRequest(w, err.Error())
		return
	}
	handlerutil.Json(w, result)
}
