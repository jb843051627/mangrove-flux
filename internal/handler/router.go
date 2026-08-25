package handler

import (
	"github.com/jb843051627/mangrove-flux/internal/service"
	"net/http"
)

func New(app *service.Lab) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health(app))
	mux.HandleFunc("/stations", stations(app))
	mux.HandleFunc("/stations/", stationDetail(app))
	mux.HandleFunc("/plots", plots(app))
	mux.HandleFunc("/chambers", chambers(app))
	mux.HandleFunc("/batches", batches(app))
	mux.HandleFunc("/batches/", batchDetail(app))
	mux.HandleFunc("/deployments", deployments(app))
	mux.HandleFunc("/deployments/", deploymentDetail(app))
	mux.HandleFunc("/readings", readings(app))
	mux.HandleFunc("/readings/", readingDetail(app))
	mux.HandleFunc("/quality/", quality(app))
	mux.HandleFunc("/alerts/", alertDetail(app))
	mux.HandleFunc("/reports/", reports(app))
	mux.HandleFunc("/calibrations", calibrations(app))
	mux.HandleFunc("/tide", tide(app))
	mux.HandleFunc("/weather", weather(app))
	mux.Handle("/ui/", static())
	return withHeaders(mux)
}

func withHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		next.ServeHTTP(w, r)
	})
}
