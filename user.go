package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/unloadingthecoding/chirpy/internal/auth"
	"github.com/unloadingthecoding/chirpy/internal/database"
)

func (cfg *apiConfig) AddUser(w http.ResponseWriter, r *http.Request) {
	type AddUserReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	e := AddUserReq{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&e)

	hashed_pw, err := auth.HashPassword(e.Password)
	if err != nil {

	}

	new_user_params := database.CreateUserParams{
		Email:          e.Email,
		HashedPassword: hashed_pw,
	}

	user, err := cfg.database.CreateUser(r.Context(), new_user_params)

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

func (cfg *apiConfig) UserLogin(w http.ResponseWriter, r *http.Request) {
	type LoginCheck struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	lc := LoginCheck{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&lc)

	user, err := cfg.database.UserLookup(r.Context(), lc.Email)
	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("Invalid email or password"))
		return
	}

	err = auth.CheckPasswordHash(lc.Password, user.HashedPassword)
	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("Invalid email or password"))
		return
	}

	type UserReply struct {
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Email     string    `json:"email"`
	}

	ur := UserReply{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}

	respondWithJSON(w, http.StatusOK, ur)
}
