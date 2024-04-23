package main

import (
	"log/slog"
	"os"
	"runtime"
)

func setupLogger() {
	appEnv := os.Getenv("APP_ENV")

	opts := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelWarn,
	}

	var handler slog.Handler = slog.NewTextHandler(os.Stdout, opts)
	if appEnv == "prod" {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	logger := slog.New(handler)
	logger = logger.With(
		slog.String("environment", appEnv),
		slog.String("level", "warn"),
		slog.String("os", runtime.GOOS),
		slog.String("arch", runtime.GOARCH),
		slog.String("version", runtime.Version()),
		slog.String("logger", "slog"),
	)

	slog.SetDefault(logger)
}
