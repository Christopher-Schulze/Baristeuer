package main

import (
	"baristeuer/internal/data"
	"baristeuer/internal/pdf"
	"baristeuer/internal/service"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func main() {
	store, err := data.NewStore("baristeuer.db")
	if err != nil {
		println("Error:", err.Error())
		return
	}
	defer store.Close()

	generator := pdf.NewGenerator("", store)
	datasvc, err := service.NewDataService("baristeuer.db")
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
