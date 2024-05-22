package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	requesterror "github.com/ItzTass/Chirpy/internal/requestError"
	responsestype "github.com/ItzTass/Chirpy/internal/responses"
)

func (cfg *apiConfig) handleChirpPost(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		requesterror.SendreqErr(w, err, "Could not decode params")
		return
	}
	if !validadeMaxLenght(w, len(params.Body)) {
		return
	}

	chirp, err := cfg.DB.CreateChirp(params.Body)
	if err != nil {
		requesterror.SendreqErr(w, err, "Internal Server Error")
		return
	}

	responsestype.RespondWithJSON(w, http.StatusCreated, chirp)
}

func (cfg *apiConfig) handleChirpsGet(w http.ResponseWriter, r *http.Request) {

	chirps, err := cfg.DB.GetChirps()
	if err != nil {
		requesterror.SendreqErr(w, err, "Internal Server Error")
		return
	}
	sort.Slice(chirps, func(i, j int) bool {
		return chirps[i].Id < chirps[j].Id
	})
	responsestype.RespondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handleChirpGetByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("chirpID"))
	if err != nil {
		requesterror.SendreqErr(w, err, "Internal Server Error")
		return
	}

	chirps, err := cfg.DB.GetChirps()
	if err != nil {
		requesterror.SendreqErr(w, err, "Internal Server Error")
		return
	}

	for _, chirp := range chirps {
		if chirp.Id == id {
			responsestype.RespondWithJSON(w, http.StatusOK, chirp)
			return
		}
	}

	responsestype.RespondWithErr(w, http.StatusNotFound, http.StatusText(http.StatusNotFound))

}
