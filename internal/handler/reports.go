package handler

import (
	"github.com/jb843051627/mangrove-flux/internal/service"
	"net/http"
	"strings"
	"time"
)

func reports(app *service.Lab) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		stationID := strings.TrimPrefix(r.URL.Path, "/reports/")
		from, _ := time.Parse(time.RFC3339, r.URL.Query().Get("from"))
		to, _ := time.Parse(time.RFC3339, r.URL.Query().Get("to"))
		if from.IsZero() {
			from = time.Now().Add(-24 * time.Hour)
		}
		if to.IsZero() {
			to = time.Now().Add(time.Hour)
		}
		if strings.HasSuffix(stationID, ".csv") {
			stationID = strings.TrimSuffix(stationID, ".csv")
			value, err := app.ExportCSV(r.Context(), stationID, from, to)
			if err != nil {
				fail(w, err)
				return
			}
			w.Header().Set("Content-Type", "text/csv")
			_, _ = w.Write(value)
			return
		}
		value, err := app.BuildReport(r.Context(), stationID, from, to)
		if err != nil {
			fail(w, err)
			return
		}
		writeJSON(w, http.StatusOK, value)
	}
}
