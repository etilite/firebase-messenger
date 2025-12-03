package http

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/etilite/firebase-messenger/internal/app"
	"github.com/stretchr/testify/require"
)

func TestBearerAuthMiddleware_Success(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		cfg := app.Config{
			APIKey: "validtoken",
		}

		mw := authMiddleware(cfg)

		handler := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Bearer validtoken")

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		require.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("invalid token", func(t *testing.T) {
		t.Parallel()

		cfg := app.Config{
			APIKey: "invalidtoken",
		}

		mw := authMiddleware(cfg)

		handler := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Authorization", "Bearer validtoken")

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		require.Equal(t, http.StatusUnauthorized, rr.Code)
	})

	t.Run("missing header", func(t *testing.T) {
		t.Parallel()

		cfg := app.Config{
			APIKey: "validtoken",
		}

		mw := authMiddleware(cfg)

		handler := mw(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		req := httptest.NewRequest("GET", "/", nil)

		rr := httptest.NewRecorder()
		handler.ServeHTTP(rr, req)

		require.Equal(t, http.StatusUnauthorized, rr.Code)
	})
}
