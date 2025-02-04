package handlerutil

import (
	"encoding/json"
	"io"
	"net/http"
)

func ReadFile(w http.ResponseWriter, r *http.Request, key string) (io.Reader, func() error, bool) {
	f, _, err := r.FormFile(key)
	if err != nil {
		BadRequest(w, "file not found")
		return nil, nil, false
	}
	return f, f.Close, true
}

func BadRequest(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(msg))
}

func Ok(w http.ResponseWriter, msg string) {
  w.WriteHeader(http.StatusOK)
  w.Write([]byte(msg))
}

func Json(w http.ResponseWriter, v interface{}) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  json.NewEncoder(w).Encode(v)
}
