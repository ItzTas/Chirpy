package main

import (
	"encoding/json"
	"net/http"
	"slices"

	"golang.org/x/crypto/bcrypt"
)

func (cfg *apiConfig) createUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	type returnVals struct {
		Email string `json:"email"`
		Id    int    `json:"id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, "Could not decode params")
		return
	}

	if slices.Contains(cfg.emails, params.Email) {
		respondWithErr(w, http.StatusBadRequest, "Email already exists")
		return
	}

	if !validadeMaxLenght(w, len(params.Email)) {
		return
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, "Could not hash password")
		return
	}

	user, err := cfg.DB.CreateUser(params.Email, string(hashed))
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, "Could not create User")
		return
	}

	cfg.emails = append(cfg.emails, params.Email)

	respondWithJSON(w, http.StatusCreated, returnVals{
		Email: user.Email,
		Id:    user.Id,
	})
}

func (cfg *apiConfig) handleLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}

}
