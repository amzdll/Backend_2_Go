package config

import "go.uber.org/config"

type ApiConfig struct {
	Stage string `yaml:"stage"`
	Host  string `yaml:"host"`
	Port  string `yaml:"port"`
}

func NewApiConfig(provider config.Provider) (*ApiConfig, error) {
	var c ApiConfig
	if err := provider.Get("application").Populate(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
