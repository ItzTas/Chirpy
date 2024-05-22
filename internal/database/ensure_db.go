package database

import (
	"errors"
	"net/http"
	"os"

	requesterror "github.com/ItzTass/Chirpy/internal/requestError"
)

func (db *DB) ensureDB() error {
	if _, err := os.Stat(db.path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			initialData := DBStructure{
				Chirps: make(map[int]Chirp),
				Users:  make(map[int]User),
			}
			if err = db.writeDB(initialData); err != nil {
				return requesterror.NewRequestErr(http.StatusInternalServerError, err.Error())
			}
		} else {
			return requesterror.NewRequestErr(http.StatusInternalServerError, err.Error())
		}
	}
	return nil
}
