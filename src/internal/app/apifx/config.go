package apifx

import "go.uber.org/config"

type Config struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func NewConfig(provider config.Provider) (*Config, error) {
	var c Config
	if err := provider.Get("application").Populate(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
