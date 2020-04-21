package profiler

import (
	"net"
	"net/http"
	"net/http/pprof"

	"go.octolab.org/ecosystem/guard/internal/config"
	"go.octolab.org/ecosystem/guard/internal/transport"
)

// New TODO issue#docs
func New(_ config.ProfilingConfig) transport.Server {
	return &server{}
}

type server struct{}

// Serve TODO issue#docs
func (*server) Serve(listener net.Listener) error {
	defer listener.Close()
	mux := &http.ServeMux{}
	mux.HandleFunc("/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/pprof/profile", pprof.Profile)
	mux.HandleFunc("/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/pprof/trace", pprof.Trace)
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	return http.Serve(listener, mux)
}
