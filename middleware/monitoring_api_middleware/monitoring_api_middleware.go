package monitoring_api_middleware

import (
	"net/http"
	"time"

	"github.com/AI-IELTS-HSE/Monitoring-Lib/service/prometheus"
)

type Middleware struct {
	MetricModels prometheus.MetricModels
}

func New(namespace string) Middleware {
	return Middleware{
		MetricModels: *prometheus.NewMetricModels(namespace),
	}
}

func (m Middleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		recorder := &statusRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(recorder, r)

		m.MetricModels.SystemMetrics.HttpMetrics.HttpStatusCodes.WithLabelValues(r.URL.Path, http.StatusText(recorder.statusCode)).Inc()

		duration := time.Since(start).Seconds()
		m.MetricModels.SystemMetrics.HttpMetrics.HttpDuration.WithLabelValues(r.URL.Path, http.StatusText(recorder.statusCode)).Observe(duration) //.HttpDuration

	})
}

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

// Override WriteHeader для записи статус-кода
func (r *statusRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}
