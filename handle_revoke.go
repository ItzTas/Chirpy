package main

import (
	"net/http"

	"github.com/ItzTass/Chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	refToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	err = cfg.DB.RevokeToken(refToken)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
