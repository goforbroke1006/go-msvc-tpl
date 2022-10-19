package panel

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"sync"
)

func InitServicePanel(port uint16) *servicePanel {
	return &servicePanel{
		port:  0,
		ready: false,
	}
}

type servicePanel struct {
	port uint16

	ready   bool
	readyMx sync.RWMutex
}

func (sp *servicePanel) SetReady(isReady bool) {
	sp.readyMx.Lock()
	defer sp.readyMx.Unlock()

	sp.ready = isReady
}

func (sp *servicePanel) isReady() bool {
	sp.readyMx.RLock()
	defer sp.readyMx.RUnlock()

	return sp.ready
}

func (sp *servicePanel) ListenAndServe(addr string) error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, "Service: %s", viper.GetString("service.name"))
	})

	mux.HandleFunc("/readyz", func(w http.ResponseWriter, req *http.Request) {
		if !sp.isReady() {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	return http.ListenAndServe(addr, mux)
}
