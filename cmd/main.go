package main

import (
	"baristeuer/internal/data"
	"baristeuer/internal/pdf"
	"baristeuer/internal/service"
	"baristeuer/internal/taxlogic"
	"flag"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"os"
)

func main() {
	dbPath := flag.String("db", "baristeuer.db", "SQLite database path")
	pdfDir := flag.String("pdfdir", "", "directory for generated PDFs")
	year := flag.Int("year", 2025, "tax year for calculations")
	logLevel := flag.String("loglevel", "", "log level (debug|info|warn|error)")
	flag.Parse()

	if *logLevel != "" {
		os.Setenv("LOG_LEVEL", *logLevel)
	}

	store, err := data.NewStore(*dbPath)
	if err != nil {
		println("Error:", err.Error())
		return
	}
	defer store.Close()

	generator := pdf.NewGenerator(*pdfDir, store, *year)
	cfg := taxlogic.ConfigForYear(*year)
	datasvc, err := service.NewDataServiceWithConfig(*dbPath, cfg)
	if err != nil {
		println("Error:", err.Error())
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
