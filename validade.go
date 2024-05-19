package main

import (
	"encoding/json"
	"net/http"
)

func handlerValidadeDecR(w http.ResponseWriter, r *http.Request) {
	type paramethers struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		Valid bool `json:"valid"`
	}
	decoder := json.NewDecoder(r.Body)
	params := paramethers{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithErr(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithErr(w, http.StatusBadRequest, "Chirp is too long")
		return
	}
	respondWithJSON(w, http.StatusOK, returnVals{
		Valid: true,
	})
}
