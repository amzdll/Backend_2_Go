package config

import (
	"go.uber.org/config"
	"go.uber.org/fx"
)

const (
	EnvLocal = "local"
	EnvDev   = "dev"
	EnvProd  = "prod"
)

type BaseConfig struct {
	Name string `yaml:"name"`
}

type Config struct {
	fx.Out
	Provider config.Provider
	Config   BaseConfig
}

// NewConfig todo: need relative path
func NewConfig() (Config, error) {
	loader, err := config.NewYAML(config.File("config/config.yaml"))
	if err != nil {
		return Config{}, err
	}
	cfg := BaseConfig{
		Name: "default",
	}
	if err = loader.Get("app").Populate(&cfg); err != nil {

		return Config{}, err
	}

	return Config{Provider: loader}, nil
}
