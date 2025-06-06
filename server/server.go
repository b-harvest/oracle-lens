package server

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"

	"bharvest.io/oracle-lens/log"
)

var GlobalState Response

func Run(listenPort int) {
	GlobalState = Response{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(GlobalState)
	})

	addr := fmt.Sprintf(":%d", listenPort)
	log.Info(fmt.Sprintf("server listening on %s", addr))

	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Error(err)
	}
}
