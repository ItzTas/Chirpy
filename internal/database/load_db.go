package database

import (
	"encoding/json"
	"net/http"
	"os"

	requesterror "github.com/ItzTass/Chirpy/internal/requestError"
)

func (db *DB) loadDB() (DBStructure, error) {
	db.mux.Lock()
	defer db.mux.Unlock()
	dat, err := os.ReadFile(db.path)
	if err != nil {
		return DBStructure{}, requesterror.NewRequestErr(http.StatusInternalServerError, err.Error())
	}

	dbStructure := DBStructure{}
	err = json.Unmarshal(dat, &dbStructure)
	if err != nil {
		return DBStructure{}, requesterror.NewRequestErr(http.StatusInternalServerError, err.Error())
	}
	return dbStructure, nil
}
