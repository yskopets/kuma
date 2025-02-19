package api_server

import (
	"errors"

	"github.com/Kong/kuma/pkg/config"
)

var _ config.Config = &ApiServerConfig{}

// API Server configuration
type ApiServerConfig struct {
	// Port of the API Server
	Port int `yaml:"port" envconfig:"kuma_api_server_port"`
	// If true, then API Server will operate in read only mode (serving GET requests)
	ReadOnly bool `yaml:"readOnly" envconfig:"kuma_api_server_read_only"`
}

func (a *ApiServerConfig) Validate() error {
	if a.Port < 0 {
		return errors.New("Port cannot be negative")
	}
	return nil
}

func DefaultApiServerConfig() *ApiServerConfig {
	return &ApiServerConfig{
		Port:     5681,
		ReadOnly: false,
	}
}
