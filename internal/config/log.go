package config

import (
	"go.uber.org/config"
)

type LogConfig struct {
	Stage string `yaml:"stage"`
	Host  string `yaml:"host"`
	Port  string `yaml:"port"`
}

func NewLogConfig(provider config.Provider) (*LogConfig, error) {
	var c LogConfig

	if err := provider.Get("application").Populate(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
