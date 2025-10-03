package middleware

import (
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	mu       sync.Mutex
	requests map[string][]time.Time
	limit    int
	interval time.Duration
}

func NewRateLimiter(limit int, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    limit,
		interval: interval,
	}
}

func (rl *RateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := r.RemoteAddr
		now := time.Now()

		rl.mu.Lock()
		times := rl.requests[ip]
		rl.mu.Unlock()

		var keep []time.Time
		for _, t := range times {
			if now.Sub(t) < rl.interval {
				keep = append(keep, t)
			}
		}
		if len(keep) >= rl.limit {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		keep = append(keep, now)
		rl.mu.Lock()
		rl.requests[ip] = keep
		rl.mu.Unlock()

		next.ServeHTTP(w, r)
	})
}
