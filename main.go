package main

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/unloadingthecoding/chirpy/handlers"
)

type payload interface {
}

func main() {

	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
	}

	mux := http.NewServeMux()
	srv := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.Handle("GET /app/", http.StripPrefix("/app/", apiCfg.middlewareMetricsInc(http.FileServer(http.Dir(".")))))
	mux.Handle("GET /assets/", http.FileServer(http.Dir("./logo.png")))
	mux.HandleFunc("GET /api/healthz", handlers.Healthzhandler)
	mux.HandleFunc("GET /admin/metrics", apiCfg.hitCount)
	mux.HandleFunc("POST /admin/reset", apiCfg.resetCount)
	mux.HandleFunc("POST /api/validate_chirp", ValidateChirp)

	err := srv.ListenAndServe()

	if err != nil {
		fmt.Printf("server reached critical state: %v", err)
	}
}
