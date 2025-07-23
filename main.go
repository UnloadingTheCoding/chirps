package main

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/unloadingthecoding/chirpy/handlers"
)

func main() {

	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
	}

	mux := http.NewServeMux()
	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.Handle("/app/", http.StripPrefix("/app/", apiCfg.middlewareMetricsInc(http.FileServer(http.Dir(".")))))
	mux.Handle("/assets", http.FileServer(http.Dir("./logo.png")))
	mux.HandleFunc("/healthz", handlers.Healthzhandler)
	mux.HandleFunc("/metrics", apiCfg.hitCount)
	mux.HandleFunc("/reset", apiCfg.resetCount)

	err := srv.ListenAndServe()

	if err != nil {
		fmt.Printf("server reached critical state: %v", err)
	}
}
