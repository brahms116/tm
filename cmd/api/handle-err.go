package main

import (
	"log"
	"net/http"
	"tm/internal/tm"
	"tm/pkg/handlerutil"
)

func (s *Server) handleErr(err error, w http.ResponseWriter) {
  if err == nil {
    return
  }

  switch e := err.(type) {
  case tm.UserErr:
    handlerutil.BadRequest(w, e.Error())
  default:
    log.Printf("Error: %v", err)
    handlerutil.ServerError(w)
  }
}
