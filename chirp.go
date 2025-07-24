package main

import (
	"encoding/json"
	"net/http"
	"slices"
	"strings"
)

type Chirp struct {
	Body string `json:"body"`
}

func ValidateChirp(w http.ResponseWriter, r *http.Request) {
	c := Chirp{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&c)

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

	type Resp struct {
		CleanedBody string `json:"cleaned_body"`
	}
	res := Resp{CleanedBody: cleaned}
	respondWithJSON(w, http.StatusOK, res)
}
