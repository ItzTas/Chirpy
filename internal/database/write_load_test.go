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

	expected := DBStructure{
		Chirps: map[int]Chirp{
			1: {
				Id:   1,
				Body: "This is the first test",
			},
			2: {
				Id:   2,
				Body: "This is the second test",
			},
			3: {
				Id:   3,
				Body: "This is the third",
			},
		},
	}

	testStructure := DBStructure{
		Chirps: map[int]Chirp{
			1: {
				Id:   1,
				Body: "This is the first test",
			},
			2: {
				Id:   2,
				Body: "This is the second test",
			},
		},
	}

	if err = db.writeDB(testStructure); err != nil {
		t.Fatalf("Expected no error to write the structure, got %v", err)
	}

	if _, err = db.CreateChirp("This is the third"); err != nil {
		t.Fatalf("Expected no error to write the chirp, got %v", err)
	}

	ts, err := db.loadDB()
	if err != nil {
		t.Fatalf("Expected no error to load the structure, got %v", err)
	}

	if len(ts.Chirps) != len(expected.Chirps) {
		t.Fatalf("Expected equal lenghts, got %v", err)
	}

	for i := range ts.Chirps {
		if ts.Chirps[i] != expected.Chirps[i] {
			t.Fatalf("Expected chirp at index %v to be %v, got %v", i, expected.Chirps[i], ts.Chirps[i])
		}
	}
}
