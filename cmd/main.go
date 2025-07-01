package main

import (
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

func main() {
	err := wails.Run(&options.App{
		Title:       "Baristeuer",
		AssetServer: &assetserver.Options{},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
