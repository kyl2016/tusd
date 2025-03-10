package cli

import (
	"net/http"

	"github.com/kyl2016/tusd"
	"github.com/kyl2016/tusd/prometheuscollector"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var MetricsOpenConnections = prometheus.NewGauge(prometheus.GaugeOpts{
	Name: "tusd_connections_open",
	Help: "Current number of open connections.",
})

var MetricsHookErrorsTotal = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "tusd_hook_errors_total",
		Help: "Total number of execution errors per hook type.",
	},
	[]string{"hooktype"},
)

func SetupMetrics(handler *tusd.Handler) {
	prometheus.MustRegister(MetricsOpenConnections)
	prometheus.MustRegister(MetricsHookErrorsTotal)
	prometheus.MustRegister(prometheuscollector.New(handler.Metrics))

	stdout.Printf("Using %s as the metrics path.\n", Flags.MetricsPath)
	http.Handle(Flags.MetricsPath, promhttp.Handler())
}
