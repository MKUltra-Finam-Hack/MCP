package middleware

import (
	"log"
	"net/http"
	"time"
)

func AuditMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		logger.Printf("[AUDIT] %s %s performed at %v", r.Method, r.URL.Path, start)
	})
}
