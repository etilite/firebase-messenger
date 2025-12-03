package app

import "os"

type Config struct {
	HostPort            string
	FirebaseCredentials string
	APIKey              string
}

func NewConfigFromEnv() Config {
	config := Config{
		HostPort:            ":3000",
		FirebaseCredentials: os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"),
		APIKey:              os.Getenv("WEBHOOK_AUTH_TOKEN"),
	}
	if httpAddr, ok := os.LookupEnv("HTTP_HOST_PORT"); ok {
		config.HostPort = httpAddr
	}

	return config
}
