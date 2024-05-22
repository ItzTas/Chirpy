package database

import (
	"net/http"

	requesterror "github.com/ItzTass/Chirpy/internal/requestError"
)

func (db *DB) GetChirps() ([]Chirp, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return []Chirp{}, requesterror.NewRequestErr(http.StatusInternalServerError, err.Error())
	}
	chirps := []Chirp{}
	for _, chirp := range dbStructure.Chirps {
		chirps = append(chirps, chirp)
	}
	return chirps, nil
}
