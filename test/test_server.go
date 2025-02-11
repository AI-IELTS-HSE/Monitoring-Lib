package main

import (
	"net/http"

	"github.com/AI-IELTS-HSE/Monitoring-Lib/middleware/monitoring_api_middleware"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	metricsMiddleware := monitoring_api_middleware.New("test")
	r := chi.NewRouter()
	r.Use(metricsMiddleware.Handle)
	r.Handle("/metrics", promhttp.Handler())
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ping"))
	})
	r.Get("/pong", func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "This endpoint always returns a 400 error", http.StatusBadRequest)
	})

	http.ListenAndServe(":8080", r)
}
