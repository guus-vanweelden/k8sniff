package main

import (
	"time"

	"github.com/golang/glog"
	"github.com/kubermatic/k8sniff/metrics"
	ping "github.com/sparrc/go-ping"
)

// SetUpPinger creates a bunch of pinger fo a list of servers
func SetUpPinger(servers []Server) {
	for _, s := range servers {
		p, err := ping.NewPinger(s.Host)
		if err != nil {
			glog.V(5).Infof("Could not setup a pinger for host=%s: %s", s.Host, err)
			continue
		}
		// TODO: make it configurable
		p.Interval = time.Second * 15
		p.Timeout = time.Second * 5

		p.OnRecv = func(pkt *ping.Packet) {
			metrics.ServerDelay(s.Host, pkt.Rtt)
		}
		go p.Run()
	}
}
