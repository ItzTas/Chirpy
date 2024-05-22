package database

import (
	"net/http"
	"sync"

	requesterror "github.com/ItzTass/Chirpy/internal/requestError"
)

func NewDB(path string) (*DB, error) {
	db := &DB{
		path: path,
		mux:  &sync.Mutex{},
	}
	if err := db.ensureDB(); err != nil {
		return nil, requesterror.NewRequestErr(http.StatusInternalServerError, err.Error())
	}
	return db, nil
}
