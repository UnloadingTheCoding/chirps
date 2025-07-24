package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/unloadingthecoding/chirpy/handlers"
	"github.com/unloadingthecoding/chirpy/internal/database"
)

type payload interface {
}

func main() {
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	db, _ := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)

	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
		database:       dbQueries,
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
	mux.HandleFunc("POST /api/users", apiCfg.AddUser)

	err := srv.ListenAndServe()

	if err != nil {
		fmt.Printf("server reached critical state: %v", err)
	}
}
