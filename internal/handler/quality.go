package handler

import (
	"github.com/jb843051627/mangrove-flux/internal/service"
	"net/http"
	"strings"
)

func quality(app *service.Lab) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stationID := strings.TrimPrefix(r.URL.Path, "/quality/")
		if r.Method != http.MethodPost {
			method(w, r, http.MethodPost)
			return
		}
		count, err := app.QualitySweep(r.Context(), stationID)
		if err != nil {
			fail(w, err)
			return
		}
		writeJSON(w, http.StatusOK, map[string]int{"evaluated": count})
	}
}
