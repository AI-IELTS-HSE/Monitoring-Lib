package prometheus

import "github.com/AI-IELTS-HSE/Monitoring-Lib/service/prometheus/product"

type MetricModels struct {
	ProductMetrics *product.ProductMetrics
}

func NewMetricModels(namespace string) *MetricModels {
	return &MetricModels{
		ProductMetrics: product.New(namespace),
	}
}
