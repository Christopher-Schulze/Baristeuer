package config

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLoadFromFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.json")
	data := `{
        "dbPath": "db.sqlite",
        "pdfDir": "./pdfs",
        "logFile": "app.log",
        "logLevel": "debug",
        "logFormat": "json",
        "taxYear": 2026,
        "formName": "Club",
        "formTaxNumber": "11/111/11111",
        "formAddress": "Street 1"
    }`
	if err := os.WriteFile(path, []byte(data), 0o644); err != nil {
		t.Fatal(err)
	}

	cfg, err := Load(path)
	if err != nil {
		t.Fatalf("Load returned error: %v", err)
	}
	if cfg.DBPath != "db.sqlite" || cfg.PDFDir != "./pdfs" || cfg.LogFile != "app.log" || cfg.LogLevel != "debug" || cfg.LogFormat != "json" ||
		cfg.TaxYear != 2026 || cfg.FormName != "Club" || cfg.FormTaxNumber != "11/111/11111" || cfg.FormAddress != "Street 1" ||
		cfg.CloudUploadURL != "" || cfg.CloudDownloadURL != "" || cfg.CloudToken != "" {
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
	if !reflect.DeepEqual(cfg, DefaultConfig) {
		t.Fatalf("expected default config, got %+v", cfg)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("config file not created: %v", err)
	}
}

func TestSaveAndVerify(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "out.json")
	expected := Config{DBPath: "db", PDFDir: "pdf", LogFile: "log", LogLevel: "info", LogFormat: "json", TaxYear: 2025, FormName: "Org", FormTaxNumber: "12/222/22222", FormAddress: "Main"}
	expected.CloudUploadURL = ""
	expected.CloudDownloadURL = ""
	expected.CloudToken = ""

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
	if !reflect.DeepEqual(got, expected) {
		t.Fatalf("written config mismatch: %+v vs %+v", got, expected)
	}
}

func TestLoadSaveError(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "missing", "config.json")

	if _, err := Load(path); err == nil {
		t.Fatal("expected error, got nil")
	}
}
