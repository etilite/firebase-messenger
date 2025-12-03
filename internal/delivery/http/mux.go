package http

import (
	"log/slog"
	"net/http"

	"github.com/etilite/firebase-messenger/internal/app"
)

func NewMux(cfg app.Config, pusher pusher) *http.ServeMux {
	h := NewPushHandler(pusher)

	mux := http.NewServeMux()

	mw := authMiddleware(cfg)
	mux.Handle("POST /webhook/send", mw(h.handlerFn()))

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{ "status": "ok" }`))
		if err != nil {
			slog.Error(err.Error())
		}
	})

	return mux
}
