package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	serverUtils "github.com/konfortes/go-server-utils/server"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func createHTTPServer(wg *sync.WaitGroup) {
	defer wg.Done()
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	mux.HandleFunc("/produceAsync", produceAsync)

	addr := fmt.Sprintf("%s:%s", host, *port)

	srv := &http.Server{
		Addr:    addr,
		Handler: mux,
	}

	go func() {
		log.Printf("Listeneing on " + addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	serverUtils.GracefulShutdown(srv)
}
