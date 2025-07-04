package service

import (
	"io"
	"log/slog"
	"os"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"
)

// NewLogger creates a slog.Logger writing to the given file.
// If logFile is empty, logs are written to stdout.
func NewLogger(logFile, level string) (*slog.Logger, io.Closer) {
	var w io.Writer = os.Stdout
	var c io.Closer
	if logFile != "" {
		lj := &lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    5, // megabytes
			MaxBackups: 3,
			MaxAge:     28, // days
			Compress:   true,
		}
		w = lj
		c = lj
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
	h := slog.NewTextHandler(w, &slog.HandlerOptions{Level: lvl})
	return slog.New(h), c
}
