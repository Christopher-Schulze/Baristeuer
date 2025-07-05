package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Config holds application configuration values.
type Config struct {
	DBPath        string `json:"dbPath"`
	PDFDir        string `json:"pdfDir"`
	LogFile       string `json:"logFile"`
	LogLevel      string `json:"logLevel"`
	TaxYear       int    `json:"taxYear"`
	FormName      string `json:"formName"`
	FormTaxNumber string `json:"formTaxNumber"`
	FormAddress   string `json:"formAddress"`
}

// DefaultConfig provides sensible defaults for a new configuration file.
var DefaultConfig = Config{
	DBPath:   "baristeuer.db",
	PDFDir:   filepath.Join(".", "internal", "data", "reports"),
	LogFile:  "baristeuer.log",
	LogLevel: "info",
}

// Load reads configuration from the given file path. If the file does not exist,
// default values are returned and no error is raised.
func Load(path string) (Config, error) {
	cfg := Config{}
	f, err := os.Open(path)
	if err != nil {
		if os.IsNotExist(err) {
			cfg = DefaultConfig
			if err := Save(path, cfg); err != nil {
				return cfg, nil
			}
			return cfg, nil
		}
		return cfg, fmt.Errorf("open config: %w", err)
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return cfg, fmt.Errorf("decode config: %w", err)
	}
	return cfg, nil
}

// Save writes the configuration to the given path in JSON format.
func Save(path string, cfg Config) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("create config: %w", err)
	}
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	if err := enc.Encode(cfg); err != nil {
		return fmt.Errorf("encode config: %w", err)
	}
	return nil
}
