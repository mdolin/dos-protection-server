// Description: A simple rate limiter that limits the number
// of requests a client can make within a given window.

package ratelimiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	requests map[string]int
	limit    int
	window   time.Duration
	mutex    sync.Mutex
}

func NewRateLimiter(limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string]int),
		limit:    limit,
		window:   window,
		mutex:    sync.Mutex{},
	}
}

func (rl *RateLimiter) AllowRequest(clientID string) bool {
	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	// Check if the client has exceeded the request limit
	count, exists := rl.requests[clientID]
	if !exists || count < rl.limit {
		// Increment the request count for the client or reset it if the window has passed
		if !exists {
			go rl.resetCount(clientID)
		}
		rl.requests[clientID]++
		return true
	}

	return false
}

func (rl *RateLimiter) resetCount(clientID string) {
	time.Sleep(rl.window)

	rl.mutex.Lock()
	defer rl.mutex.Unlock()

	delete(rl.requests, clientID)
}
