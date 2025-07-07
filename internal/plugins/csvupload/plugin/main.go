package main

import (
	"baristeuer/internal/plugins"
	"baristeuer/internal/plugins/csvupload"
)

// New is the symbol loaded by the application.
func New() plugins.Plugin {
	return csvupload.New()
}
