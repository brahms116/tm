package main

import (
	"net/http"
	"tm/pkg/handlerutil"
)

func (s *Server) importCsv(w http.ResponseWriter, r *http.Request) {
	f, close, ok := handlerutil.ReadFile(w, r, "file")
	if !ok {
		return
	}
	defer close()

	result, err := s.tm.ImportCsv(r.Context(), f)
	if err != nil {
    s.handleErr(err, w)
		return
	}
	handlerutil.Json(w, result)
}

