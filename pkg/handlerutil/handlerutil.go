package handlerutil

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func ReadOptionalQueryInt(w http.ResponseWriter, r *http.Request, key string, defaultValue int) (int, bool) {
	str := r.URL.Query().Get(key)
	if str == "" {
		return defaultValue, true
	}
	i, err := strconv.Atoi(str)
	if err != nil {
		BadRequest(w, "invalid integer")
		return 0, false
	}
	return i, true
}

func ReadOptionalQueryString(r *http.Request, key string, defaultValue string) string {
	str := r.URL.Query().Get(key)
	if str == "" {
		return defaultValue
	}
	return str
}

func ReadQueryTime(
	w http.ResponseWriter,
	r *http.Request,
	key string,
	format string,
) (time.Time, bool) {
	str := r.URL.Query().Get(key)
	if str == "" {
		BadRequest(w, fmt.Sprintf("missing %s query parameter", key))
		return time.Time{}, false
	}
	t, err := time.Parse(format, str)
	if err != nil {
		BadRequest(w, fmt.Sprintf("invalid %s query parameter, expected format: %s", key, format))
		return time.Time{}, false
	}
	return t, true
}

func ReadFile(w http.ResponseWriter, r *http.Request, key string) (io.Reader, func() error, bool) {
	f, _, err := r.FormFile(key)
	if err != nil {
		BadRequest(w, "file not found")
		return nil, nil, false
	}
	return f, f.Close, true
}

func BadRequest(w http.ResponseWriter, msg string) {
  http.Error(w, msg, http.StatusBadRequest)
}

func ServerError(w http.ResponseWriter) {
  http.Error(w, "Internal server error", http.StatusInternalServerError)
}

func Unauthorized(w http.ResponseWriter, msg string) {
  http.Error(w, msg, http.StatusUnauthorized)
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

func Text(w http.ResponseWriter, v string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(v))
}

func BodyJson[T any](w http.ResponseWriter, r *http.Request) (T, bool) {
	var t T
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&t)
  if err != nil {
    BadRequest(w, "Malformed json")
    return t, false
  }
	return t, true
}
