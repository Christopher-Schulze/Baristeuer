package main

import (
	"baristeuer/internal/data"
	"baristeuer/internal/pdf"
	"baristeuer/internal/service"
	"flag"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func main() {
	dbPath := flag.String("db", "baristeuer.db", "SQLite database path")
	pdfDir := flag.String("pdfdir", "", "directory for generated PDFs")
	flag.Parse()

	store, err := data.NewStore(*dbPath)
	if err != nil {
		println("Error:", err.Error())
		return
	}
	defer store.Close()

	generator := pdf.NewGenerator(*pdfDir, store)
	datasvc, err := service.NewDataService(*dbPath)
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
