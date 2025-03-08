package main

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestToken(t *testing.T) {
	app := InitApp()

	// Request without token
	req := httptest.NewRequest("GET", "/check-token", nil)
	req.Header.Set("Authorization", "")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != fiber.StatusUnauthorized {
		t.Errorf("Expected status %d, got %d", fiber.StatusUnauthorized, resp.StatusCode)
	}

	// Request with token
	req = httptest.NewRequest("GET", "/check-token", nil)
	req.Header.Set("Authorization", "Bearer mock-token")
	resp, err = app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status %d, got %d", fiber.StatusOK, resp.StatusCode)
	}
}