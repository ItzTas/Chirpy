package main

import (
	"encoding/json"
	"net/http"
	"sort"

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
	chirp, err := db.CreateChirp(params.Body)
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, chirp)
}

func handleChirpsGet(w http.ResponseWriter, r *http.Request) {
	db, err := database.NewDB(database_path)
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, err.Error())
	}

	chirps, err := db.GetChirps()
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, err.Error())
	}
	sort.Slice(chirps, func(i, j int) bool {
		return chirps[i].Id < chirps[j].Id
	})
	respondWithJSON(w, http.StatusOK, chirps)
}
