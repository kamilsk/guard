package monitor

import (
	"expvar"
	"net"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/kamilsk/guard/pkg/config"
	"github.com/kamilsk/guard/pkg/transport"
)

// New TODO issue#docs
func New(_ config.MonitoringConfig) transport.Server {
	return &server{}
}

type server struct{}

// Serve TODO issue#docs
func (*server) Serve(listener net.Listener) error {
	defer listener.Close()
	mux := &http.ServeMux{}
	mux.Handle("/metrics", promhttp.Handler())
	mux.Handle("/vars", expvar.Handler())
	return http.Serve(listener, mux)
}
