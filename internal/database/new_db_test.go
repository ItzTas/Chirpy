package database

import (
	"errors"
	"os"
	"testing"
)

func TestNewDb(t *testing.T) {
	const path = "./database.json"
	os.Remove(path)

	if _, err := NewDB(path); err != nil {
		t.Fatalf("NewDB failed: %s", err)
	}

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		t.Fatalf("Expected file %s to be created, but it does not exist", path)
	}

	os.Remove(path)
}
