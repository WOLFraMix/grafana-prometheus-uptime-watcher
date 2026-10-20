package main

import (
	"log"
	"net/http"

	"github.com/WOLFraMix/grafana-prometheus-uptime-watcher/internal/config"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	cfg := config.New()

	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("ok"))
		if err != nil {
			log.Println(err)
		}
	})

	h := promhttp.Handler()
	mux.Handle("/metrics", h)

	srv := &http.Server{Addr: cfg.Port, Handler: mux}

	log.Println("server is listening on", cfg.Port)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
