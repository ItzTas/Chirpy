package database

import (
	"encoding/json"
	"net/http"
	"os"

	requesterror "github.com/ItzTass/Chirpy/internal/requestError"
)

func (db *DB) writeDB(dbStructure DBStructure) error {
	db.mux.Lock()
	defer db.mux.Unlock()
	dat, err := json.Marshal(dbStructure)
	if err != nil {
		return requesterror.NewRequestErr(http.StatusInternalServerError, err.Error())
	}

	if err = os.WriteFile(db.path, dat, 0644); err != nil {
		return requesterror.NewRequestErr(http.StatusInternalServerError, err.Error())
	}
	return nil
}
