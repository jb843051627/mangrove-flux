package handler

import (
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/service"
	"net/http"
	"strings"
)

func stations(app *service.Lab) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var value model.FieldStation
			if err := decode(r, &value); err != nil {
				fail(w, err)
				return
			}
			if err := app.CreateStation(r.Context(), value); err != nil {
				fail(w, err)
				return
			}
			writeJSON(w, http.StatusCreated, value)
		case http.MethodGet:
			values, err := app.ListStations(r.Context())
			if err != nil {
				fail(w, err)
				return
			}
			writeJSON(w, http.StatusOK, values)
		default:
			method(w, r, http.MethodGet)
		}
	}
}
func stationDetail(app *service.Lab) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/stations/")
		if r.Method == http.MethodDelete {
			if err := app.DisableStation(r.Context(), id); err != nil {
				fail(w, err)
				return
			}
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if !method(w, r, http.MethodGet) {
			return
		}
		value, err := app.GetStation(r.Context(), id)
		if err != nil {
			fail(w, err)
			return
		}
		writeJSON(w, http.StatusOK, value)
	}
}
