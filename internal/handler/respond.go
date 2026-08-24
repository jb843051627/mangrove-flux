package handler

import (
	"encoding/json"
	"errors"
	"github.com/jb843051627/mangrove-flux/internal/model"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}
func decode(r *http.Request, value any) error { return json.NewDecoder(r.Body).Decode(value) }
func fail(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	if errors.Is(err, model.ErrNotFound) {
		status = http.StatusNotFound
	}
	if errors.Is(err, model.ErrInvalid) || errors.Is(err, model.ErrQuality) {
		status = http.StatusBadRequest
	}
	if errors.Is(err, model.ErrConflict) {
		status = http.StatusConflict
	}
	if errors.Is(err, model.ErrCancelled) {
		status = http.StatusRequestTimeout
	}
	writeJSON(w, status, map[string]string{"error": err.Error()})
}
func method(w http.ResponseWriter, r *http.Request, expected string) bool {
	if r.Method != expected {
		w.Header().Set("Allow", expected)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false
	}
	return true
}
