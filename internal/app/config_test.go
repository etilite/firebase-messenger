package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfigFromEnv(t *testing.T) {
	t.Run("default http address", func(t *testing.T) {
		cfg := NewConfigFromEnv()

		assert.Equal(t, ":3000", cfg.HostPort)
	})

	t.Run("http address from env", func(t *testing.T) {
		t.Setenv("HTTP_HOST_PORT", "test:7777")
		cfg := NewConfigFromEnv()

		assert.Equal(t, "test:7777", cfg.HostPort)
	})

	t.Run("webhook api key", func(t *testing.T) {
		t.Setenv("WEBHOOK_AUTH_TOKEN", "test-api-key")
		cfg := NewConfigFromEnv()

		assert.Equal(t, "test-api-key", cfg.APIKey)
	})

	t.Run("firebase credentials file", func(t *testing.T) {
		t.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "some/file.json")
		cfg := NewConfigFromEnv()

		assert.Equal(t, "some/file.json", cfg.FirebaseCredentials)
	})
}
