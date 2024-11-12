package main

import "time"

type RateLimiter struct {
	Requests []time.Time
	Rate     time.Duration
	Limit    int
}

func NewRateLimiter(n int, t time.Duration) *RateLimiter {
	return &RateLimiter{Requests: make([]time.Time, n), Rate: t}
}

func (rl *RateLimiter) Allow() bool {
	timestamp := time.Now()

	if rl.Full() {
		for i, req := range rl.Requests {
			if time.Now().Sub(req) > rl.Rate {
				rl.Requests = rl.Requests[:i]
				rl.Requests = append(rl.Requests, timestamp)
				return true
			}
		}
		return false
	} else {
		rl.Requests = append(rl.Requests, timestamp)
		return true
	}
}

func (rl *RateLimiter) Full() bool {
	return len(rl.Requests) == rl.Limit
}

func (rl *RateLimiter) Empty() bool {
	return len(rl.Requests) == 0
}
