package handler

import (
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/service"
	"net/http"
)

func calibrations(app *service.Lab) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !method(w, r, http.MethodPost) {
			return
		}
		var value model.CalibrationProfile
		if err := decode(r, &value); err != nil {
			fail(w, err)
			return
		}
		if err := app.CreateCalibration(r.Context(), value); err != nil {
			fail(w, err)
			return
		}
		writeJSON(w, http.StatusCreated, value)
	}
}
func tide(app *service.Lab) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !method(w, r, http.MethodPost) {
			return
		}
		var value model.TideObservation
		if err := decode(r, &value); err != nil {
			fail(w, err)
			return
		}
		if err := app.RecordTide(r.Context(), value); err != nil {
			fail(w, err)
			return
		}
		writeJSON(w, http.StatusCreated, value)
	}
}
func weather(app *service.Lab) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !method(w, r, http.MethodPost) {
			return
		}
		var value model.WeatherObservation
		if err := decode(r, &value); err != nil {
			fail(w, err)
			return
		}
		if err := app.RecordWeather(r.Context(), value); err != nil {
			fail(w, err)
			return
		}
		writeJSON(w, http.StatusCreated, value)
	}
}
