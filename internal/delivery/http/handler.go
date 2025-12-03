package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/etilite/firebase-messenger/internal/model"
)

type pusher interface {
	Push(ctx context.Context, request model.SendRequest) (model.SendResponse, error)
}

type PushHandler struct {
	pusher pusher
}

func NewPushHandler(pusher pusher) *PushHandler {
	return &PushHandler{pusher: pusher}
}

func (h *PushHandler) handlerFn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sendRequest := model.SendRequest{}
		if err := json.NewDecoder(r.Body).Decode(&sendRequest); err != nil {
			err = fmt.Errorf("failed to decode JSON: %v", err)
			slog.Error("handler: bad request", "error", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		response, err := h.pusher.Push(r.Context(), sendRequest)
		if err != nil {
			err = fmt.Errorf("failed to push: %v", err)
			slog.Error("handler: internal error", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(response); err != nil {
			slog.Error("handler: failed to write response", "error", err)
		}
	}
}
