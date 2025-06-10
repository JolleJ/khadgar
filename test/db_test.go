package test

import (
	"jollej/db-scout/internal/infrastructure"
	"testing"
)

func TestDBInitialization(t *testing.T) {
	_, err := infrastructure.InitDb()

	if err != nil {
		t.Fatalf("Failed to initialze the database %v", err)
	}
}
