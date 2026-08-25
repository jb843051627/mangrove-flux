package handler

import (
	"github.com/jb843051627/mangrove-flux/internal/model"
	"github.com/jb843051627/mangrove-flux/internal/service"
	"net/http"
	"strings"
)

func deployments(app *service.Lab) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !method(w, r, http.MethodPost) {
			return
		}
		var value model.Deployment
		if err := decode(r, &value); err != nil {
			fail(w, err)
			return
		}
		if err := app.PlanDeployment(r.Context(), value); err != nil {
			fail(w, err)
			return
		}
		writeJSON(w, http.StatusCreated, value)
	}
}
func deploymentDetail(app *service.Lab) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/deployments/")
		switch {
		case strings.HasSuffix(id, "/start"):
			id = strings.TrimSuffix(id, "/start")
			if err := app.StartDeployment(r.Context(), id); err != nil {
				fail(w, err)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		case strings.HasSuffix(id, "/close"):
			id = strings.TrimSuffix(id, "/close")
			if err := app.CloseDeployment(r.Context(), id); err != nil {
				fail(w, err)
				return
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			value, err := app.GetDeployment(r.Context(), id)
			if err != nil {
				fail(w, err)
				return
			}
			writeJSON(w, http.StatusOK, value)
		}
	}
}
