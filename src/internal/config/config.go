package config

import (
	"go.uber.org/config"
	"go.uber.org/fx"
)

type Config struct {
	Name string `yaml:"name"`
}

type ResultConfig struct {
	fx.Out
	Provider config.Provider
	Config   Config
}

func NewConfig() (ResultConfig, error) {
	loader, err := config.NewYAML(config.File("config.yaml"))
	if err != nil {
		return ResultConfig{}, err
	}

	config := Config{
		Name: "default",
	}

	if err = loader.Get("app").Populate(&config); err != nil {
		return ResultConfig{}, err
	}

	return ResultConfig{Provider: loader}, nil
}
