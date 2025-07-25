package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/google/uuid"
	"github.com/unloadingthecoding/chirpy/internal/database"
)

func (cfg *apiConfig) GenerateChirp(w http.ResponseWriter, r *http.Request) {
	req := database.CreateChirpParams{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	c, err := cfg.database.CreateChirp(r.Context(), req)

	if err != nil {
		respondWithError(w, 400, "Something went wrong")
	}

	if len(c.Body) > 140 {
		respondWithError(w, 400, "Chirp is too long")
	}

	clean := []string{}
	for _, word := range strings.Split(c.Body, " ") {
		if slices.Contains([]string{"kerfuffle", "sharbert", "fornax"}, strings.ToLower(word)) {
			clean = append(clean, "****")
			continue
		}
		clean = append(clean, word)
	}
	cleaned := strings.Join(clean, " ")

	c.Body = cleaned

	respondWithJSON(w, 201, c)
}

func (cfg *apiConfig) GetAllChirps(w http.ResponseWriter, r *http.Request) {
	chirps, err := cfg.database.GetAllChirps(r.Context())
	if err != nil {
		log.Printf("unable to retrieve all chirps: %s", err)
	}

	respondWithJSON(w, 200, chirps)

}

func (cfg *apiConfig) GetChirp(w http.ResponseWriter, r *http.Request) {
	chirpID := r.PathValue("chirpID")

	parsedID, err := uuid.Parse(chirpID)
	if err != nil {
		log.Printf("malformed chirp id: %s", err)
	}

	chirp, err := cfg.database.GetChirp(r.Context(), parsedID)

	if err != nil {
		msg := fmt.Sprintf("chirp not found")
		respondWithError(w, 404, msg)
		return
	}

	respondWithJSON(w, 200, chirp)
}
