package handler

import (
	"github.com/jb843051627/mangrove-flux/internal/service"
	"net/http"
)

func health(app *service.Lab) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !method(w, r, http.MethodGet) {
			return
		}
		value, err := app.Health(r.Context())
		if err != nil {
			fail(w, err)
			return
		}
		writeJSON(w, http.StatusOK, value)
	}
}
