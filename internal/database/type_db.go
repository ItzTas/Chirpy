package database

import (
	"sync"
)

type DB struct {
	path string
	mux  *sync.Mutex
}

type DBStructure struct {
	Chirps map[int]Chirp `json:"chirps"`
}

type Chirp struct {
	id   int
	body string
}
