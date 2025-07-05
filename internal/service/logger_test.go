package service

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNewLogger_FormatSwitch(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "log.txt")

	logger, closer := NewLogger(path, "info", "text")
	if closer == nil {
		t.Fatalf("expected closer")
	}
	logger.Info("msg")
	closer.Close()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "level=INFO") {
		t.Fatalf("expected text log, got %s", data)
	}

	logger, closer = NewLogger(path, "info", "json")
	logger.Info("msg")
	closer.Close()
	data, err = os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "\"level\":\"INFO\"") {
		t.Fatalf("expected json log, got %s", data)
	}
}
