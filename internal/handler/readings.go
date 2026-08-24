package handler

import (
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/service"
	"net/http"
)

func readings(app *service.Lab) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !method(w, r, http.MethodPost) {
			return
		}
		if r.URL.Query().Get("batch") == "1" {
			var values []model.FluxReading
			if err := decode(r, &values); err != nil {
				fail(w, err)
				return
			}
			if err := app.BatchIngest(r.Context(), values); err != nil {
				fail(w, err)
				return
			}
			writeJSON(w, http.StatusAccepted, map[string]int{"count": len(values)})
			return
		}
		var value model.FluxReading
		if err := decode(r, &value); err != nil {
			fail(w, err)
			return
		}
		if err := app.IngestReading(r.Context(), value); err != nil {
			fail(w, err)
			return
		}
		writeJSON(w, http.StatusCreated, value)
	}
}
func readingDetail(app *service.Lab) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Path[len("/readings/"):]
		if r.URL.Query().Get("deployment_id") != "" {
			values, err := app.RecentReadings(r.Context(), r.URL.Query().Get("deployment_id"))
			if err != nil {
				fail(w, err)
				return
			}
			writeJSON(w, http.StatusOK, values)
			return
		}
		value, err := app.EvaluateReading(r.Context(), id)
		if err != nil {
			fail(w, err)
			return
		}
		writeJSON(w, http.StatusOK, value)
	}
}
