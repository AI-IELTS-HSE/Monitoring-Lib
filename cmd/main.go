package main
import (
	"github.com/AI-IELTS-HSE/Monitoring-Lib/middleware/monitoring_api_middleware"
	"github.com/go-chi/chi/v5"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"

)

func main(){

	metricsMiddleware:= monitoring_api_middleware.New("test")
	r := chi.NewRouter()
	r.Use(metricsMiddleware.Handle)
	r.Handle("/metrics", promhttp.Handler())
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ping"))
	})

	http.ListenAndServe(":8080", r)
}