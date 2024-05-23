package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ItzTass/Chirpy/internal/auth"
)

func (cfg *apiConfig) handlerChirpsDelete(w http.ResponseWriter, r *http.Request) {
	chirpIDString := r.PathValue("chirpID")
	chirpID, err := strconv.Atoi(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "Could not convert the chirp id to integer")
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusForbidden, err.Error())
		return
	}

	userIDstr, err := auth.ValidateJWT(token, cfg.jwtSecret)
	if err != nil {
		respondWithError(w, http.StatusForbidden, err.Error())
		return
	}
	userID, err := strconv.Atoi(userIDstr)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "Could not convert the user id to integer")
		return
	}

	chirp, err := cfg.DB.GetChirp(chirpID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not find the chirp in the database")
		return
	}

	if chirp.AuthorID != userID {
		fmt.Println(userID)
		fmt.Println(chirp.AuthorID)
		respondWithError(w, http.StatusForbidden, "Can't delete the chirpy of another author")
		return
	}

	if err = cfg.DB.DeleteChirp(chirpID); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
