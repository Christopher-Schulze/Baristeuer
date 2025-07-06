package main

import (
	"baristeuer/internal/plugins"
	"baristeuer/internal/plugins/logging"
)

// New is the symbol loaded by the application.
func New() plugins.Plugin {
	return logging.New()
}
