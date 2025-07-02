package main

import (
	"net/http"
	"tm/pkg/handlerutil"
)

func (s *Server) health(w http.ResponseWriter, r *http.Request) {
	handlerutil.Ok(w, "OK")
}
