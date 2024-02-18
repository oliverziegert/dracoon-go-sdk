package core

import (
	"golang.org/x/oauth2"
	log "log/slog"
	"net/http"
	"os"
)

type Config struct {
	Token2 oauth2.Token

	// OAuth2 access token
	Token string

	// OAuth2 refresh token
	RefreshToken string

	// Logging target for verbose SDK logging
	Logger *log.Logger

	// Logging level for SDK generated logs
	LogLevel log.Level

	// Logging format for SDK generated logs
	LogHandler log.Handler

	// DRACOON target domain
	Domain string

	// No need to set -- for testing only
	Client *http.Client

	// RetryMaxAttempts specifies the maximum number attempts an API client
	// will call an operation that fails with a retryable error.
	RetryMaxAttempts int
}

func NewConfig() *Config {
	return &Config{}
}

func NewDefaultConfig() *Config {
	cfg := NewConfig()
	logLevel := log.LevelVar{}
	logLevel.Set(log.LevelWarn)
	cfg.LogHandler = log.NewTextHandler(os.Stdout, &log.HandlerOptions{
		AddSource: true,
		Level:     &logLevel,
	})
	cfg.Logger = log.New(cfg.LogHandler)
	cfg.Domain = defaultDomain
	cfg.RetryMaxAttempts = defaultRetryMaxAttempts
	return cfg
}
