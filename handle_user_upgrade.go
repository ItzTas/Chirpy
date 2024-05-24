package main

import (
	"encoding/json"
	"net/http"

	"github.com/ItzTass/Chirpy/internal/auth"
)

const (
	userUpgradedEvent = "user.upgraded"
)

func (cfg *apiConfig) handlerUsersUpgrade(w http.ResponseWriter, r *http.Request) {
	type paramethers struct {
		Event string `json:"event"`
		Data  struct {
			UserID int `json:"user_id"`
		} `json:"data"`
	}

	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid key")
		return
	}

	if apiKey != cfg.polkaSecret {
		respondWithError(w, http.StatusUnauthorized, "Not authorized")
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := paramethers{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not decode paramethers")
		return
	}

	if params.Event != userUpgradedEvent {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	err = cfg.DB.UpgradeUserToRed(params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
