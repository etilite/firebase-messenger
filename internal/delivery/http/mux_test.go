package http

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/etilite/firebase-messenger/internal/app"
	"github.com/etilite/firebase-messenger/internal/model"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
)

func TestNewMux(t *testing.T) {
	t.Parallel()
	tests := map[string]struct {
		method string
		path   string
	}{
		"/push/send": {
			method: http.MethodPost,
			path:   "/webhook/send",
		},
		"/health": {
			method: http.MethodGet,
			path:   "/health",
			// todo check health body?
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			mc := minimock.NewController(t)
			pusherMock := NewPusherMock(mc)
			pusherMock.PushMock.Optional().Set(func(ctx context.Context, request model.SendRequest) (s1 model.SendResponse, err error) {
				return model.SendResponse{}, nil
			})
			mux := NewMux(app.Config{APIKey: "some-token"}, pusherMock)

			requestBody := strings.NewReader(`{}`)
			request := httptest.NewRequest(tt.method, tt.path, requestBody)
			request.Header.Set("Authorization", "Bearer some-token")
			response := httptest.NewRecorder()

			mux.ServeHTTP(response, request)

			assert.Equal(t, http.StatusOK, response.Code)
		})
	}
}
