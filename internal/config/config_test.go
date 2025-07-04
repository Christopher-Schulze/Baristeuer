package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestLoadFromFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")
	data := `{
        "dbPath": "db.sqlite",
        "pdfDir": "./pdfs",
        "logFile": "app.log",
        "logLevel": "debug"
    }`
	if err := os.WriteFile(path, []byte(data), 0o644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}
	if cfg.DBPath != "db.sqlite" || cfg.PDFDir != "./pdfs" || cfg.LogFile != "app.log" || cfg.LogLevel != "debug" {
		t.Fatalf("unexpected config: %+v", cfg)
	}
}

func TestLoadMissingFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")

	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}
	if cfg != (Config{}) {
		t.Fatalf("expected zero config, got %+v", cfg)
	}
}

func TestSaveAndVerify(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.json")
	expected := Config{DBPath: "db", PDFDir: "pdf", LogFile: "log", LogLevel: "info"}

	if err := Save(path, expected); err != nil {
		t.Fatalf("Save returned error: %v", err)
	}

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("reading file failed: %v", err)
	}

	var got Config
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}
	if got != expected {
		t.Fatalf("written config mismatch: %+v vs %+v", got, expected)
	}
}
