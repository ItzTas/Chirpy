package database

import (
	"sync"
)

func NewDB(path string) (*DB, error) {
	db := &DB{
		path: path,
		mux:  &sync.Mutex{},
	}
	if err := db.ensureDB(); err != nil {
		return nil, err
	}
	return db, nil

}
