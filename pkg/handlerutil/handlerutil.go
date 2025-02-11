package handlerutil

import (
	"encoding/json"
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

func ReadOptionalQueryDate(
	w http.ResponseWriter,
	r *http.Request,
	key string,
	defaultTime time.Time,
) (time.Time, bool) {
	str := r.URL.Query().Get(key)
	if str == "" {
		return defaultTime, true
	}
	t, err := time.Parse("2006-01-02", str)
	if err != nil {
		BadRequest(w, "invalid date format")
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
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(msg))
}

func Unauthorized(w http.ResponseWriter, msg string) {
	w.WriteHeader(http.StatusUnauthorized)
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
