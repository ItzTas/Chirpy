package database

import (
	"net/http"

	requesterror "github.com/ItzTass/Chirpy/internal/requestError"
)

func (db *DB) CreateChirp(body string) (Chirp, error) {
	if err := db.ensureDB(); err != nil {
		return Chirp{}, requesterror.NewRequestErr(http.StatusInternalServerError, err.Error())
	}
	dbstru, err := db.loadDB()
	if err != nil {
		return Chirp{}, requesterror.NewRequestErr(http.StatusInternalServerError, err.Error())
	}
	newID := len(dbstru.Chirps) + 1
	newC := Chirp{
		Id:   newID,
		Body: body,
	}
	dbstru.Chirps[newID] = newC
	if err = db.writeDB(dbstru); err != nil {
		return Chirp{}, requesterror.NewRequestErr(http.StatusInternalServerError, err.Error())
	}
	return newC, nil
}
