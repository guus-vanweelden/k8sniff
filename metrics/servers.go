package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	serversDelayHistVec = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "servers_ping_histogram_seconds",
		Help: "Ping duration to the other servers.",
	},
		[]string{"server"})
)

// ServerDelay gather the duration to a specific server
func ServerDelay(server string, d time.Duration) {
	serversDelayHistVec.With(prometheus.Labels{"server": server}).Observe(d.Seconds())
}
