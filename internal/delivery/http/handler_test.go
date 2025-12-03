package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/etilite/firebase-messenger/internal/model"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlerFn(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mc := minimock.NewController(t)
		pusherMock := NewPusherMock(mc)
		pusherMock.PushMock.Expect(minimock.AnyContext, model.SendRequest{}).Return(model.SendResponse{
			SuccessCount: 1,
			FailureCount: 1,
			Responses: []model.TokenSendResult{{
				Success:   true,
				MessageID: "1",
				Error:     nil,
			}},
		}, nil)

		server := NewPushHandler(pusherMock)
		requestBody := strings.NewReader(`{}`)
		request, _ := http.NewRequest(http.MethodPost, "/", requestBody)
		response := httptest.NewRecorder()
		handleFunc := server.handlerFn()

		handleFunc(response, request)
		defer request.Body.Close()

		require.Equal(t, http.StatusOK, response.Code)
		assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
		want :=
			`{"successCount":1,"failureCount":1,"responses":[{"success":true,"messageId":"1"}]}
`
		require.Equal(t, want, response.Body.String())
	})

	t.Run("bad request", func(t *testing.T) {
		t.Parallel()

		mc := minimock.NewController(t)
		pusherMock := NewPusherMock(mc)
		server := NewPushHandler(pusherMock)

		requestBody := strings.NewReader("")
		request := httptest.NewRequest(http.MethodPost, "/", requestBody)
		response := httptest.NewRecorder()
		handleFunc := server.handlerFn()

		handleFunc(response, request)
		defer request.Body.Close()

		require.Equal(t, http.StatusBadRequest, response.Code)
		require.Contains(t, response.Body.String(), "failed to decode JSON: EOF")
	})

	t.Run("internal server error", func(t *testing.T) {
		t.Parallel()

		mc := minimock.NewController(t)
		pusherMock := NewPusherMock(mc)
		pusherMock.PushMock.Expect(minimock.AnyContext, model.SendRequest{}).Return(model.SendResponse{}, assert.AnError)

		server := NewPushHandler(pusherMock)

		requestBody := strings.NewReader(`{"size": 32, "content": "test"}`)
		request := httptest.NewRequest(http.MethodPost, "/", requestBody)
		response := httptest.NewRecorder()
		handleFunc := server.handlerFn()

		handleFunc(response, request)

		require.Equal(t, http.StatusInternalServerError, response.Code)
		require.Contains(t, response.Body.String(), "failed to push: "+assert.AnError.Error())
	})
}
