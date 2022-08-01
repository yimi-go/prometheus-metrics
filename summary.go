package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/yimi-go/metrics"
)

type summary struct {
	sv     *prometheus.SummaryVec
	labels map[string]string
}

// NewSummary new a prometheus summary and returns Observer
func NewSummary(sv *prometheus.SummaryVec) metrics.Observer {
	return &summary{
		sv: sv,
	}
}

func (s *summary) With(labels map[string]string) metrics.Observer {
	merged := make(map[string]string, len(s.labels)+len(labels))
	for k, v := range s.labels {
		merged[k] = v
	}
	for k, v := range labels {
		merged[k] = v
	}
	return &summary{
		sv:     s.sv,
		labels: merged,
	}
}

func (s *summary) Observe(value float64) {
	s.sv.With(s.labels).Observe(value)
}
