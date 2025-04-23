package middleware

import (
	"github.com/jie10/crud-auth-go/internal/config"
	"net/http"
	"strings"
)

func SecurityMiddleware(next http.Handler, cfg *config.Config) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if len(cfg.CORSAllowedOrigins) == 0 || contains(cfg.CORSAllowedOrigins, "*") || contains(cfg.CORSAllowedOrigins, origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("X-Frame-Options", "DENY")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

		// Content Security Policy
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data:;")

		// Handle OPTIONS for CORS preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
		}

		// Prevent MIME sniffing
		if r.Header.Get("Content-Type") != "" {
			contentType := r.Header.Get("Content-Type")
			if !strings.HasPrefix(contentType, "application/json") {
				http.Error(w, "Unsupported Media Type", http.StatusUnsupportedMediaType)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
	return mux
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
