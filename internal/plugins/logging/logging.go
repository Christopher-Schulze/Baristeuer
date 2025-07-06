package logging

import (
	"baristeuer/internal/plugins"
	"baristeuer/internal/service"
)

type Plugin struct{}

func New() plugins.Plugin {
	return &Plugin{}
}

func (p *Plugin) Init(ds *service.DataService) error {
	service.Logger().Info("logging plugin initialized")
	return nil
}
