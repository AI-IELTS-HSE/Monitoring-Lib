package product

import "github.com/prometheus/client_golang/prometheus"

type HttpMetrics struct {
	HttpStatusCodes *prometheus.CounterVec
	HttpDuration    *prometheus.HistogramVec
}

func InitHttpMetrics(namespace, subsystem string) *HttpMetrics {
	hm := &HttpMetrics{
		HttpStatusCodes: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "http_status_codes_total",
				Help:      "Total number of HTTP status codes returned by each handler.",
			},
			[]string{"handler", "code"},
		),
		HttpDuration: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: namespace,
				Subsystem: subsystem,
				Name:      "http_response_duration_ms",
				Help:      "Histogram of response durations for HTTP requests.",
				Buckets:   []float64{0.01, 0.02, 0.03, 0.05, 0.075, 0.1, 0.15, 0.2, 0.3, 0.4, 0.5, 0.75, 1, 2, 5},
			},
			[]string{"handler", "status_code"},
		),
	}

	// Регистрация метрики в Prometheus
	prometheus.MustRegister(hm.HttpStatusCodes)
	prometheus.MustRegister(hm.HttpDuration)
	return hm
}

// func init() {
// 	// Регистрация метрики в Prometheus
// 	prometheus.MustRegister(HttpStatusCodes)
// 	prometheus.MustRegister(HttpDuration)
// }
