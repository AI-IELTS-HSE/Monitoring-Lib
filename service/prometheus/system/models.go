package system

type SystemMetrics struct {
	HttpMetrics *HttpMetrics
}

func New(namespace string) *SystemMetrics {
	pm := &SystemMetrics{
		HttpMetrics: InitHttpMetrics(namespace, SubSystemSystem),
	}
	return pm
}
