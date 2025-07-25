package main

import (
	"sync/atomic"

	"github.com/unloadingthecoding/chirpy/internal/database"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	database       *database.Queries
	platform       string
}
