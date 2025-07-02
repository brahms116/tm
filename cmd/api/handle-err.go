package main

import (
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
    handlerutil.ServerError(w)
  }
}
