package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/yimi-go/metrics"
)

type counter struct {
	cv     *prometheus.CounterVec
	labels map[string]string
}

// NewCounter new a prometheus counter and returns Counter.
func NewCounter(cv *prometheus.CounterVec) metrics.Counter {
	return &counter{
		cv: cv,
	}
}

func (c *counter) With(labels map[string]string) metrics.Counter {
	merged := make(map[string]string, len(c.labels)+len(labels))
	for k, v := range c.labels {
		merged[k] = v
	}
	for k, v := range labels {
		merged[k] = v
	}
	return &counter{
		cv:     c.cv,
		labels: merged,
	}
}

func (c *counter) Inc() {
	c.cv.With(c.labels).Inc()
}

func (c *counter) Add(delta float64) {
	c.cv.With(c.labels).Add(delta)
}
