package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (cfg *apiConfig) AddUser(w http.ResponseWriter, r *http.Request) {
	type email struct {
		Email string `json:"email"`
	}

	e := email{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&e)

	user, err := cfg.database.CreateUser(r.Context(), e.Email)

	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Unable to generate user [db]: %s", err))
	}

	respondWithJSON(w, 201, user)
}
