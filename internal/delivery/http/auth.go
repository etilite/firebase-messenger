package http

import (
	"net/http"
	"strings"

	"github.com/etilite/firebase-messenger/internal/app"
)

func authMiddleware(config app.Config) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if !strings.HasPrefix(authHeader, "Bearer ") {
				http.Error(w, "missing or invalid Authorization header", http.StatusUnauthorized)
				return
			}

			token := strings.TrimPrefix(authHeader, "Bearer ")
			if token == "" {
				http.Error(w, "empty bearer token", http.StatusUnauthorized)
				return
			}

			if token != config.APIKey {
				http.Error(w, "invalid token", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
