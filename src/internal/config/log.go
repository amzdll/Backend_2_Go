package config

import (
	"go.uber.org/config"
)

type LogConfig struct {
	Status string `yaml:"status"`
}

func NewLogConfig(provider config.Provider) (*LogConfig, error) {
	var c LogConfig

	if err := provider.Get("env").Populate(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
