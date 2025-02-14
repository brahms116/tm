package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"tm/internal/cfg"
	"tm/internal/tm"
	"tm/pkg/handlerutil"
)

type Server struct {
	cfg cfg.Cfg
	tm  tm.TM
}

func New(cfg cfg.Cfg, tm tm.TM) *Server {
	return &Server{tm: tm, cfg: cfg}
}

func (s *Server) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", s.health)
	mux.Handle("POST /import", s.authMiddleware(http.HandlerFunc(s.importIngCsv)))

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		return fmt.Errorf("error starting server: %w", err)
	}
	return nil
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	handlerutil.Ok(w, "OK")
}

// GET /report?month=2021-01
func (s *Server) report(w http.ResponseWriter, r *http.Request) {
	dateMonth, ok := handlerutil.ReadQueryTime(w, r, "month", "2006-01")
	if !ok {
		return
	}

	result, err := s.tm.Report(r.Context(), dateMonth)
	if err != nil {
		handlerutil.BadRequest(w, err.Error())
		return
	}
	handlerutil.Json(w, result)
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

func (s *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("Authorization")
		apiKey = strings.ReplaceAll(apiKey, "Bearer ", "")
		if apiKey != s.cfg.ApiKey {
			handlerutil.Unauthorized(w, "Invalid api key")
			return
		}
		next.ServeHTTP(w, r)
	})
}
