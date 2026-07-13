package logger

import (
	"log/slog"
	"os"
)

func NewJSONLogger(env string) *slog.Logger {
	level := slog.LevelInfo

	if env == "local" || env == "development" {
		level = slog.LevelDebug
	}

	handler := slog.NewJSONHandler(
		os.Stdout,
		&slog.HandlerOptions{
			Level: level,
		},
	)

	return slog.New(handler)
}
