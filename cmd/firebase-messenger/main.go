package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/etilite/firebase-messenger/internal/app"
	httpserver "github.com/etilite/firebase-messenger/internal/delivery/http"
	"github.com/etilite/firebase-messenger/internal/firebase"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer done()

	cfg := app.NewConfigFromEnv()

	sender, err := firebase.New(ctx, cfg)
	if err != nil {
		slog.Error("error initializing firebase sender", "error", err)
		os.Exit(1)
	}

	server := httpserver.NewServer(cfg.HostPort, httpserver.NewMux(cfg, sender))

	if err := server.Run(ctx); err != nil {
		slog.Error("unable to start app", "error", err)
		os.Exit(1)
	}
}
