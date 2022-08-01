package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/yimi-go/metrics"
)

type histogram struct {
	hv     *prometheus.HistogramVec
	labels map[string]string
}

// NewHistogram new a prometheus histogram and returns Observer.
func NewHistogram(hv *prometheus.HistogramVec) metrics.Observer {
	return &histogram{
		hv: hv,
	}
}

func (h *histogram) With(labels map[string]string) metrics.Observer {
	merged := make(map[string]string, len(h.labels)+len(labels))
	for k, v := range h.labels {
		merged[k] = v
	}
	for k, v := range labels {
		merged[k] = v
	}
	return &histogram{
		hv:     h.hv,
		labels: merged,
	}
}

func (h *histogram) Observe(value float64) {
	h.hv.With(h.labels).Observe(value)
}
