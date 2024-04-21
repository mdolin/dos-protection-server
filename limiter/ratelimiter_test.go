package ratelimiter

import (
	"testing"
	"time"
)

func TestNewRateLimiter(t *testing.T) {
	limit := 5
	window := 10 * time.Second

	rl := NewRateLimiter(limit, window)

	if rl == nil {
		t.Error("NewRateLimiter returned nil")
		return
	}

	if rl.limit != limit {
		t.Errorf("Expected limit to be %d, got %d", limit, rl.limit)
	}

	if rl.window != window {
		t.Errorf("Expected window to be %v, got %v", window, rl.window)
	}

	if rl.requests == nil {
		t.Error("Expected requests map to be initialized")
	}
}

func TestAllowRequest(t *testing.T) {
	limit := 5
	window := 5 * time.Second

	rl := NewRateLimiter(limit, window)

	clientID := "client1"

	// Allow requests within the limit
	for i := 0; i < limit; i++ {
		if !rl.AllowRequest(clientID) {
			t.Errorf("Expected AllowRequest to return true for client %s, iteration %d", clientID, i+1)
		}
	}

	// Deny requests beyond the limit
	if rl.AllowRequest(clientID) {
		t.Errorf("Expected AllowRequest to return false for client %s after exceeding limit", clientID)
	}

	// Wait for the window to reset
	time.Sleep(window)

	// Allow requests again after the window has reset
	if !rl.AllowRequest(clientID) {
		t.Errorf("Expected AllowRequest to return true for client %s after window reset", clientID)
	}
}
