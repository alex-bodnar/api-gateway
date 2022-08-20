package config

import (
	"fmt"
	"time"

	"proxy/pkg/config"
	"proxy/pkg/errs"
	"proxy/pkg/log"
)

const (
	// DefaultPath - default path for config.
	DefaultPath = "./cmd/config.yaml"
)

type (
	// Config defines the properties of the application configuration.
	Config struct {
		Logger               log.Config           `yaml:"logger"`
		Delivery             Delivery             `yaml:"delivery"`
		AuthorizationService AuthorizationService `yaml:"authorization-service"`
		Proxy                []Proxy              `yaml:"proxy"`
	}

	// Delivery defines API server configuration.
	Delivery struct {
		HTTPServer HTTPServer `yaml:"http-server"`
	}

	// HTTPServer defines HTTP section of the API server configuration.
	HTTPServer struct {
		LogRequests        bool          `yaml:"log-requests"`
		ListenAddress      string        `yaml:"listen-address"`
		ReadTimeout        time.Duration `yaml:"read-timeout"`
		WriteTimeout       time.Duration `yaml:"write-timeout"`
		BodySizeLimitBytes int           `yaml:"body-size-limit"`
		GracefulTimeout    int           `yaml:"graceful-timeout"`
	}

	// Proxy is all about proxying routes.
	Proxy struct {
		Enabled bool         `yaml:"enabled"`
		Group   string       `yaml:"group"`
		Routes  []ProxyRoute `yaml:"routes"`
	}

	// ProxyRoute defines a proxy destination.
	ProxyRoute struct {
		In                 string `yaml:"in"`
		To                 string `yaml:"to"`
		Method             string `yaml:"method"`
		CheckAuthorization bool   `yaml:"check-authorization"`
	}

	// AuthorizationService defines authorization service configuration.
	AuthorizationService struct {
		URL string `yaml:"url"`
	}
)

// New loads and validates all configuration data, returns filled Cfg - configuration data model.
func New(appName, cfgFilePath string) (*Config, error) {
	cfg := new(Config)

	if cfgErr := cfg.loadFromFile(cfgFilePath); cfgErr != nil {
		return nil, fmt.Errorf("config loader: %s", cfgErr)
	}

	return cfg.valid()
}

// loadFromFile loads configuration from file.
func (c *Config) loadFromFile(configPath string) error {
	if err := config.LoadFromFile(configPath, c); err != nil {
		return err
	}

	return nil
}

// valid validates configuration data.
func (c *Config) valid() (*Config, error) {
	if errorsList := c.Validate(); len(errorsList) != 0 {
		return nil, &errs.FieldsValidation{Errors: errorsList}
	}

	return c, nil
}
