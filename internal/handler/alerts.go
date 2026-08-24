package handler

import (
	"github.com/jb843051627/mangrove-flux/internal/service"
	"net/http"
	"strings"
)

func alertDetail(app *service.Lab) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/alerts/")
		if strings.HasSuffix(id, "/review") {
			id = strings.TrimSuffix(id, "/review")
			if err := app.ReviewAlert(r.Context(), id); err != nil {
				fail(w, err)
				return
			}
			w.WriteHeader(http.StatusNoContent)
			return
		}
		if strings.HasSuffix(id, "/clear") {
			id = strings.TrimSuffix(id, "/clear")
			if err := app.ClearAlert(r.Context(), id); err != nil {
				fail(w, err)
				return
			}
			w.WriteHeader(http.StatusNoContent)
			return
		}
		values, err := app.ListAlerts(r.Context(), id)
		if err != nil {
			fail(w, err)
			return
		}
		writeJSON(w, http.StatusOK, values)
	}
}
