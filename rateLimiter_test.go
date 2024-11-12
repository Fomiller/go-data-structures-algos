package main

// import (
// 	"testing"
// 	"time"
// )

// commented because it has time.Sleep
// func TestRateLimiter(t *testing.T) {
// 	// Create a rate limiter allowing 3 requests every 1 second
// 	rl := RateLimiter{
// 		Limit: 3,
// 		Rate:  time.Second,
// 	}
//
// 	// Allow the first three requests
// 	for i := 0; i < 3; i++ {
// 		if !rl.Allow() {
// 			t.Errorf("Request %d should be allowed, but it was denied", i+1)
// 		}
// 	}
//
// 	// Wait for the window to reset
// 	time.Sleep(time.Second)
//
// 	for i := 0; i < 3; i++ {
// 		if !rl.Allow() {
// 			t.Errorf("Request %d should be allowed, but it was denied", i+1)
// 		}
// 	}
//
// 	// The fourth request should be denied
// 	if rl.Allow() {
// 		t.Error("Fourth request should be denied, but it was allowed")
// 	}
//
// 	// Wait for the window to reset
// 	time.Sleep(time.Second)
//
// 	// After the window reset, the first request should be allowed again
// 	if !rl.Allow() {
// 		t.Error("Request should be allowed after time window reset, but it was denied")
// 	}
// }
