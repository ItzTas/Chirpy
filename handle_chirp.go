package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
)

func (cfg *apiConfig) handleChirpPost(w http.ResponseWriter, r *http.Request) {
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

	chirp, err := cfg.DB.CreateChirp(params.Body)
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, chirp)
}

func (cfg *apiConfig) handleChirpsGet(w http.ResponseWriter, r *http.Request) {

	chirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}
	sort.Slice(chirps, func(i, j int) bool {
		return chirps[i].Id < chirps[j].Id
	})
	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handleChirpGetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("chirpID"))
	if err != nil {
		respondWithErr(w, http.StatusNotFound, err.Error())
		return
	}

	chirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, err.Error())
		return
	}

	for _, chirp := range chirps {
		if chirp.Id == id {
			respondWithJSON(w, http.StatusOK, chirp)
			return
		}
	}

	respondWithErr(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))

}
