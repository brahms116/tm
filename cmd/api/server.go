package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"tm/internal/cfg"
	"tm/internal/tm"
	"tm/pkg/contracts"
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
	mux.Handle("GET /report", s.authMiddleware(http.HandlerFunc(s.report)))
	mux.Handle("GET /is-authenticated", s.authMiddleware(http.HandlerFunc(s.isAuthenicated)))
	mux.Handle("POST /report-period", s.authMiddleware(http.HandlerFunc(s.reportPeriod)))
	mux.Handle("POST /report-timeline", s.authMiddleware(http.HandlerFunc(s.reportTimeline)))

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		return fmt.Errorf("error starting server: %w", err)
	}
	return nil
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	handlerutil.Ok(w, "OK")
}

func (s *Server) isAuthenicated(w http.ResponseWriter, r *http.Request) {
	handlerutil.Ok(w, "OK")
}

func (s *Server) reportTimeline(w http.ResponseWriter, r *http.Request) {
	reqBody, ok := handlerutil.BodyJson[contracts.TimelineRequest](w, r)
	if !ok {
		return
	}

	start, err := time.Parse(time.RFC3339, reqBody.StartDate)
	if err != nil {
		handlerutil.BadRequest(w, "Invalid start date")
		return
	}

	end, err := time.Parse(time.RFC3339, reqBody.EndDate)
	if err != nil {
		handlerutil.BadRequest(w, "Invalid end date")
		return
	}

	res, err := s.tm.ReportTimeline(
		r.Context(),
		start,
		end,
	)

	if err != nil {
		log.Println(err.Error())
		handlerutil.ServerError(w)
		return
	}

	handlerutil.Json(w, res)
	return

}

func (s *Server) reportPeriod(w http.ResponseWriter, r *http.Request) {
	reqBody, ok := handlerutil.BodyJson[contracts.ReportRequest](w, r)
	if !ok {
		return
	}

	start, err := time.Parse(time.RFC3339, reqBody.StartDate)
	if err != nil {
		handlerutil.BadRequest(w, "Invalid start date")
		return
	}

	end, err := time.Parse(time.RFC3339, reqBody.EndDate)
	if err != nil {
		handlerutil.BadRequest(w, "Invalid end date")
		return
	}

	res, err := s.tm.ReportPeriod(
		r.Context(),
		start,
		end,
		reqBody.U100,
	)

	if err != nil {
		log.Println(err.Error())
		handlerutil.ServerError(w)
		return
	}

	handlerutil.Json(w, res)
	return
}

// GET /report?month=2021-01&format=text
func (s *Server) report(w http.ResponseWriter, r *http.Request) {
	dateMonth, ok := handlerutil.ReadQueryTime(w, r, "month", "2006-01")
	if !ok {
		return
	}

	format := handlerutil.ReadOptionalQueryString(r, "format", "json")
	if format == "text" {

		result, err := s.tm.ReportText(r.Context(), dateMonth)
		if err != nil {
			handlerutil.BadRequest(w, err.Error())
			return
		}
		handlerutil.Text(w, result)
	} else {
		result, err := s.tm.Report(r.Context(), dateMonth)
		if err != nil {
			handlerutil.BadRequest(w, err.Error())
			return
		}
		handlerutil.Json(w, result)
	}
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
