package service

import (
	"io"
	"log/slog"
	"os"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"
)

var logLevelVar = new(slog.LevelVar)

func parseLevel(level string) slog.Level {
	switch strings.ToLower(level) {
	case "debug":
		return slog.LevelDebug
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	}
	return slog.LevelInfo
}

// NewLogger creates a slog.Logger writing to the given file.
// If logFile is empty, logs are written to stdout.
func NewLogger(logFile, level string) (*slog.Logger, io.Closer) {
	var w io.Writer = os.Stdout
	var c io.Closer
	if logFile != "" {
		lj := &lumberjack.Logger{
			Filename:   logFile,
			MaxSize:    5,
			MaxBackups: 3,
			MaxAge:     28,
			Compress:   true,
		}
		w = lj
		c = lj
	}
	logLevelVar.Set(parseLevel(level))
	h := slog.NewTextHandler(w, &slog.HandlerOptions{Level: logLevelVar})
	return slog.New(h), c
}

// SetLogLevel updates the global log level used by all loggers.
func SetLogLevel(level string) {
	logLevelVar.Set(parseLevel(level))
}
