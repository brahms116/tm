package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
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

//go:embed dist
var dist embed.FS

func (s *Server) Start() error {
	distFs, err := fs.Sub(dist, "dist")
	if err != nil {
		return fmt.Errorf("error getting dist dir: %w", err)
	}
	httpFs := http.FS(distFs)

	fserver := http.FileServer(httpFs)
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", s.health)
	mux.Handle("POST /import", s.applyMiddlewares(http.HandlerFunc(s.importIngCsv)))
	mux.Handle("POST /is-authenticated", s.applyMiddlewares(http.HandlerFunc(s.isAuthenicated)))
	mux.Handle("POST /report-period", s.applyMiddlewares(http.HandlerFunc(s.reportPeriod)))
	mux.Handle("POST /report-timeline", s.applyMiddlewares(http.HandlerFunc(s.reportTimeline)))
	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upath := r.URL.Path
		if !strings.HasPrefix(upath, "/") {
			upath = "/" + upath
			r.URL.Path = upath
		}
		upath = path.Clean(upath)
		f, err := httpFs.Open(upath)
		if os.IsNotExist(err) {
			r.URL.Path = "/"
			fserver.ServeHTTP(w, r)
			return
		}
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		f.Close()
		fserver.ServeHTTP(w, r)
		return
	}))

	handler := s.corsMiddleware(mux)

	err = http.ListenAndServe(":8081", handler)
	if err != nil {
		return fmt.Errorf("error starting server: %w", err)
	}
	return nil
}

func (s *Server) applyMiddlewares(h http.Handler) http.Handler {
	return s.authMiddleware(h)
}

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	handlerutil.Ok(w, "OK")
}

func (s *Server) isAuthenicated(w http.ResponseWriter, r *http.Request) {
	handlerutil.Json(w, true)
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

func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
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
