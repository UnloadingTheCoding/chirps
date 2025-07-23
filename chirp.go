package main

import (
	"encoding/json"
	"net/http"
)

type Chirp struct {
	Body string `json:"body"`
}

func ValidateChirp(w http.ResponseWriter, r *http.Request) {
	c := Chirp{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&c)

	resBody := make(map[string]interface{})

	if err != nil {
		resBody["error"] = "Something went wrong"
		res, _ := json.Marshal(resBody)
		w.WriteHeader(400)
		w.Write(res)
	}

	if len(c.Body) > 140 {
		resBody["error"] = "Chirp is too long"
		res, _ := json.Marshal(resBody)
		w.WriteHeader(400)
		w.Write(res)
	}

	resBody["valid"] = true
	res, _ := json.Marshal(resBody)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
