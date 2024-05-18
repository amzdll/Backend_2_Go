package db

import "go.uber.org/config"

type Config struct {
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	SslMode  string `yaml:"ssl_mode"`
	Dsn      string `yaml:"dsn"`
}

func NewConfig(provider config.Provider) (*Config, error) {
	var c Config
	if err := provider.Get("db").Populate(&c); err != nil {
		return nil, err
	}
	return &c, nil
}
