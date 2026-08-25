package handler

import (
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/service"
	"net/http"
	"strings"
)

func batches(app *service.Lab) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !method(w, r, http.MethodPost) {
			if r.Method == http.MethodGet {
				values, err := app.ListBatches(r.Context())
				if err != nil {
					fail(w, err)
					return
				}
				writeJSON(w, http.StatusOK, values)
			}
			return
		}
		var value model.SurveyBatch
		if err := decode(r, &value); err != nil {
			fail(w, err)
			return
		}
		if err := app.OpenBatch(r.Context(), value); err != nil {
			fail(w, err)
			return
		}
		writeJSON(w, http.StatusCreated, value)
	}
}
func batchDetail(app *service.Lab) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/batches/")
		if r.Method == http.MethodPost && strings.HasSuffix(id, "/close") {
			id = strings.TrimSuffix(id, "/close")
			if err := app.CloseBatch(r.Context(), id); err != nil {
				fail(w, err)
				return
			}
			w.WriteHeader(http.StatusNoContent)
			return
		}
		value, err := app.GetBatch(r.Context(), id)
		if err != nil {
			fail(w, err)
			return
		}
		writeJSON(w, http.StatusOK, value)
	}
}
