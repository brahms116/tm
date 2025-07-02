package main

import (
	"net/http"
	"tm/pkg/handlerutil"
)

// Super contrived authentication system.
// Currently there is only an api key, which the auth middleware checks.
// If the request passes the middleware then we return that the user is authenticated
// When I have more time, this will removed in favor of a proper authentication system with signed JWTs
func (s *Server) isAuthenicated(w http.ResponseWriter, r *http.Request) {
	handlerutil.Json(w, true)
}
