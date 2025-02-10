package product

type ProductMetrics struct {
	HttpMetrics HttpMetrics
}

var ProdMetrics ProductMetrics

func (pm *ProductMetrics) New(namespace string) *ProductMetrics {
	pm.HttpMetrics.InitHttpMetrics(namespace, SubSystemProduct)
	return pm
}
