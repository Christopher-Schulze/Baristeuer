// Package plugins defines the minimal interface for runtime extensions.
package plugins

import "baristeuer/internal/service"

// Plugin defines the interface for optional runtime extensions.
type Plugin interface {
	// Init is called once after the application has started and provides
	// access to the DataService.
	Init(*service.DataService) error
}
