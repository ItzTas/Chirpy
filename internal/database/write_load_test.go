package database

import (
	"os"
	"testing"
)

func TestNewDBstruc(t *testing.T) {
	path := "./test_database.json"
	defer os.Remove(path)

	db, err := NewDB(path)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	testStructure := DBStructure{
		Chirps: map[int]Chirp{
			1: {
				Id:   1,
				Body: "This is the first test",
			},
		},
	}
	err = db.writeDB(testStructure)
	if err != nil {
		t.Fatalf("Expected no error to write the structure, got %v", err)
	}

	ts, err := db.loadDB()
	if err != nil {
		t.Fatalf("Expected no error to load the structure, got %v", err)
	}

	if ts.Chirps[1] != testStructure.Chirps[1] {
		t.Fatalf("Expected the structures to be equal, got %v", err)
	}
}
