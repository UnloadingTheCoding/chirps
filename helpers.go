package main

import (
	"encoding/json"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {

	type errorResponse struct {
		Error string `json:"error"`
	}
	e := errorResponse{}
	e.Error = msg

	respondWithJSON(w, code, e.Error)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(payload)
	w.WriteHeader(code)
	w.Write([]byte(res))
}
