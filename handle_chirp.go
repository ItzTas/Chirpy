package main

import (
	"encoding/json"
	"net/http"

	"github.com/ItzTass/Chirpy/internal/database"
)

func handleChirpPost(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	if !validadeMaxLenght(w, len(params.Body)) {
		return
	}

	db, err := database.NewDB(database_path)
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	cleaned := validadeProfane(params.Body)
	chirp, err := db.CreateChirp(cleaned)
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, chirp)
}
