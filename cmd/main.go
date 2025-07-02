package main

import (
	"baristeuer/internal/pdf"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func main() {
	generator := pdf.NewGenerator("")
	err := wails.Run(&options.App{
		Title:       "Baristeuer",
		AssetServer: &assetserver.Options{},
		Bind:        []interface{}{generator},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
