package prometheus

import "github.com/AI-IELTS-HSE/Monitoring-Lib/service/prometheus/system"

type MetricModels struct {
	SystemMetrics *system.SystemMetrics
}

func NewMetricModels(namespace string) *MetricModels {
	return &MetricModels{
		SystemMetrics: system.New(namespace),
	}
}
