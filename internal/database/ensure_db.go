package database

import (
	"errors"
	"os"
)

func (db *DB) ensureDB() error {
	if _, err := os.Stat(db.path); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			initialData := DBStructure{
				Chirps: make(map[int]Chirp),
			}
			if err = db.writeDB(initialData); err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}
