package main

import (
	"baristeuer/internal/config"
	"baristeuer/internal/data"
	"baristeuer/internal/pdf"
	"baristeuer/internal/service"
	"flag"
	"fmt"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func main() {
	cfgPath := flag.String("config", "config.json", "configuration file")
	dbPath := flag.String("db", "", "SQLite database path")
	pdfDir := flag.String("pdfdir", "", "directory for generated PDFs")
	logFile := flag.String("logfile", "", "log file path")
	logLevel := flag.String("loglevel", "", "log level")
	flag.Parse()

	cfg, err := config.Load(*cfgPath)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	if *dbPath != "" {
		cfg.DBPath = *dbPath
	}
	if *pdfDir != "" {
		cfg.PDFDir = *pdfDir
	}
	if *logFile != "" {
		cfg.LogFile = *logFile
	}
	if *logLevel != "" {
		cfg.LogLevel = *logLevel
	}

	if cfg.DBPath == "" {
		cfg.DBPath = "baristeuer.db"
	}

	store, err := data.NewStore(cfg.DBPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer store.Close()

	logger, logCloser := service.NewLogger(cfg.LogFile, cfg.LogLevel)
	generator := pdf.NewGenerator(cfg.PDFDir, store)
	datasvc, err := service.NewDataService(cfg.DBPath, logger, logCloser)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer datasvc.Close()

	err = wails.Run(&options.App{
		Title:       "Baristeuer",
		AssetServer: &assetserver.Options{},
		Bind:        []interface{}{generator, datasvc},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
