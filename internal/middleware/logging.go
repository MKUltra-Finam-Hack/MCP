package middleware

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		logger.Printf("Completed %s %s in %v", r.Method, r.URL.Path, time.Since(start))
	})
}
