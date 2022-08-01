package prometheus

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/yimi-go/metrics"
)

type gauge struct {
	gv     *prometheus.GaugeVec
	labels map[string]string
}

// NewGauge new a prometheus gauge and returns Gauge.
func NewGauge(gv *prometheus.GaugeVec) metrics.Gauge {
	return &gauge{
		gv: gv,
	}
}

func (g *gauge) With(labels map[string]string) metrics.Gauge {
	merged := make(map[string]string, len(g.labels)+len(labels))
	for k, v := range g.labels {
		merged[k] = v
	}
	for k, v := range labels {
		merged[k] = v
	}
	return &gauge{
		gv:     g.gv,
		labels: merged,
	}
}

func (g *gauge) Set(value float64) {
	g.gv.With(g.labels).Set(value)
}

func (g *gauge) Add(delta float64) {
	g.gv.With(g.labels).Add(delta)
}

func (g *gauge) Sub(delta float64) {
	g.gv.With(g.labels).Sub(delta)
}
