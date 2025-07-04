package service

import (
	"io"
	"log/slog"
	"os"
	"strings"
)

// NewLogger creates a slog.Logger writing to the given file. The format
// parameter determines the output style: "json" or "text". If logFile is
// empty, logs are written to stdout.
func NewLogger(logFile, level, format string) (*slog.Logger, io.Closer) {
	var w io.Writer = os.Stdout
	var c io.Closer
	if logFile != "" {
		f, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
		if err == nil {
			w = f
			c = f
		}
	}
	if format == "" {
		format = os.Getenv("BARISTEUER_LOGFORMAT")
	}
	if format == "" {
		format = "text"
	}
	lvl := slog.LevelInfo
	switch strings.ToLower(level) {
	case "debug":
		lvl = slog.LevelDebug
	case "warn":
		lvl = slog.LevelWarn
	case "error":
		lvl = slog.LevelError
	}
	opts := &slog.HandlerOptions{Level: lvl}
	var h slog.Handler
	switch strings.ToLower(format) {
	case "json":
		h = slog.NewJSONHandler(w, opts)
	default:
		h = slog.NewTextHandler(w, opts)
	}
	return slog.New(h), c
}
