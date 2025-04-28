package main

import (
	"os"
	"testing"
)

func TestSetupEnvironment(t *testing.T) {
	t.Setenv("APP_URL", "http://test")
	t.Setenv("PORT", "8080")
	t.Setenv("REDIS_URL", "redis://localhost:6379")

	err := setupEnvironment()
	if err != nil {
		t.Fatalf("setupEnvironment() failed: %v", err)
	}

	if os.Getenv("PORT") != "8080" {
		t.Fatalf("setupEnvironment() failed: wrong environment loaded")
	}
}
