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
	"tm/internal/cfg"
	"tm/internal/tm"
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
	mux.Handle("POST /import", s.applyMiddlewares(http.HandlerFunc(s.importCsv)))

  // Routes are post atm, as haven't figured built in custom field annotations for in our contracts generation tool.
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


