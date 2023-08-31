package config

import (
	"log"
	"time"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/fx"
)

type Configuration struct {
	HTTPServer
	Database
	Cache
}

type HTTPServer struct {
	ServerAddress string        `envconfig:"SERVE_ADDRESS" default:"0.0.0.0"`
	IdleTimeout   time.Duration `envconfig:"HTTP_SERVER_IDLE_TIMEOUT" default:"60s"`
	Port          int           `envconfig:"PORT" default:"8080"`
	ReadTimeout   time.Duration `envconfig:"HTTP_SERVER_READ_TIMEOUT" default:"1s"`
	WriteTimeout  time.Duration `envconfig:"HTTP_SERVER_WRITE_TIMEOUT" default:"2s"`
}

func Load() (*Configuration, error) {
	var cfg Configuration
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// NewConfig creates and load config
func NewConfig() *Configuration {
	cfg, err := Load()
	if err != nil {
		log.Printf("Can't load the configuration. Error: %s", err.Error())
	}

	return cfg
}

// Module
var ConfigModule = fx.Options(
	fx.Provide(NewConfig),
)
