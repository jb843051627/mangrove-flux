package handler

import (
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/service"
	"net/http"
)

func chambers(app *service.Lab) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var value model.Chamber
			if err := decode(r, &value); err != nil {
				fail(w, err)
				return
			}
			if err := app.RegisterChamber(r.Context(), value); err != nil {
				fail(w, err)
				return
			}
			writeJSON(w, http.StatusCreated, value)
			return
		}
		stationID := r.URL.Query().Get("station_id")
		values, err := app.ListChambers(r.Context(), stationID)
		if err != nil {
			fail(w, err)
			return
		}
		writeJSON(w, http.StatusOK, values)
	}
}
