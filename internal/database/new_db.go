package database

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

func NewDB(path string) (*DB, error) {
	if _, err := os.Stat(path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			initialData := DBStructure{
				Chirps: make(map[int]Chirp),
			}
			dat, err := json.Marshal(initialData)
			if err != nil {
				return nil, err
			}
			err = os.WriteFile(path, dat, 0644)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}
	db := &DB{
		path: path,
		mux:  &sync.Mutex{},
	}
	return db, nil

}
