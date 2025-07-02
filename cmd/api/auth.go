package main

import (
	"net/http"
	"tm/pkg/handlerutil"
)

// Super contrived authentication system.
// Currently there is only an api key, which the auth middleware checks.
// If the request passes the middleware then we return that the user is authenticated
// When I have more time, this will return a properly formed JWT.
func (s *Server) isAuthenicated(w http.ResponseWriter, r *http.Request) {
	handlerutil.Json(w, true)
}
