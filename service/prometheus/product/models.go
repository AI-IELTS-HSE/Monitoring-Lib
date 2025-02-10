package product

type ProductMetrics struct {
	HttpMetrics *HttpMetrics
}

func New(namespace string) *ProductMetrics {
	pm := &ProductMetrics{
		HttpMetrics: InitHttpMetrics(namespace, SubSystemProduct),
	}
	return pm
}
