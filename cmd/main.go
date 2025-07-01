package main

import (
    "github.com/wailsapp/wails/v2"
    "github.com/wailsapp/wails/v2/pkg/options"

    "baristeuer/internal/backend"
)

func main() {
    app := backend.NewBackend()
    if err := wails.Run(&options.App{Bind: []interface{}{app}}); err != nil {
        println("Error:", err.Error())
    }
}
