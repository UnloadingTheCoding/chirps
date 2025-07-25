package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/unloadingthecoding/chirpy/internal/auth"
	"github.com/unloadingthecoding/chirpy/internal/database"
)

func (cfg *apiConfig) AddUser(w http.ResponseWriter, r *http.Request) {
	e := database.CreateUserParams{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&e)

	hashed_pw, err := auth.HashPassword(e.HashedPassword)
	if err != nil {

	}

	e.HashedPassword = hashed_pw

	user, err := cfg.database.CreateUser(r.Context(), e)

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Unable to generate user [db]: %s", err))
	}

	respondWithJSON(w, 201, user)
}

func (cfg *apiConfig) DeleteAllUsers(w http.ResponseWriter, r *http.Request) {
	if cfg.platform != "dev" {
		respondWithError(w, http.StatusForbidden, "UNAUTHORIZED ACCESS")
	}

	err := cfg.database.DeleteAllUsers(r.Context())
	if err != nil {
		log.Printf("Unable to delete users: %s", err)
	}
	w.WriteHeader(http.StatusOK)
}
