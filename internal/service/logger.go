package service

import (
	"io"
	"log/slog"
	"os"
	"strings"

	"gopkg.in/natefinch/lumberjack.v2"
)

func init() {
	stdLogger = slog.New(slog.NewTextHandler(logWriter, &slog.HandlerOptions{Level: logLevelVar}))
}

var (
	logLevelVar           = new(slog.LevelVar)
	logWriter   io.Writer = os.Stdout
	stdLogger   *slog.Logger
)

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
func NewLogger(logFile, level, format string) (*slog.Logger, io.Closer) {
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
	logWriter = w
	logLevelVar.Set(parseLevel(level))
	var h slog.Handler
	if strings.ToLower(format) == "json" {
		h = slog.NewJSONHandler(logWriter, &slog.HandlerOptions{Level: logLevelVar})
	} else {
		h = slog.NewTextHandler(logWriter, &slog.HandlerOptions{Level: logLevelVar})
	}
	stdLogger = slog.New(h)
	return stdLogger, c
}

// SetLogLevel updates the global log level used by all loggers.
func SetLogLevel(level string) {
	logLevelVar.Set(parseLevel(level))
}

// SetLogFormat replaces the handler of the global logger with the given format.
func SetLogFormat(format string) {
	var h slog.Handler
	if strings.ToLower(format) == "json" {
		h = slog.NewJSONHandler(logWriter, &slog.HandlerOptions{Level: logLevelVar})
	} else {
		h = slog.NewTextHandler(logWriter, &slog.HandlerOptions{Level: logLevelVar})
	}
	stdLogger = slog.New(h)
}

// Logger returns the currently configured global logger.
func Logger() *slog.Logger {
	return stdLogger
}
