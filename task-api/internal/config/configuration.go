package config

import (
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	DB   DB     `yaml:"db"`
	Port int    `envconfig:"APP_PORT" yaml:"port"`
	Env  string `envconfig:"APP_ENV" yaml:"env"`
}

type DB struct {
	Host     string `envconfig:"APP_DB_HOST" yaml:"host"`
	User     string `envconfig:"APP_DB_USER" yaml:"user"`
	Pass     string `envconfig:"APP_DB_PASS" yaml:"pass"`
	Name     string `envconfig:"APP_DB_NAME" yaml:"name"`
	Port     int    `envconfig:"APP_DB_PORT" yaml:"port"`
	MaxConns int    `envconfig:"APP_DB_MAX_CONN" yaml:"maxConn"`
}

func NewConfig(filePath string) (*AppConfig, error) {
	config, err := loadFromYaml(filePath)
	if err != nil {
		return nil, err
	}

	if err := loadFromEnvVars(config); err != nil {
		return nil, err
	}

	return config, nil
}

func loadFromYaml(filePath string) (*AppConfig, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var config AppConfig

	if err = yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func loadFromEnvVars(config *AppConfig) error {
	return envconfig.Process("APP", config)
}
