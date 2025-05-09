package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type ClientConfig struct {
	Postgres     PostgresConfig `yaml:"postgres"`
	ServerConfig ServerConfig   `yaml:"client-server"`
}

type ServerConfig struct {
	Port         int  `yaml:"port"`
	IsProduction bool `yaml:"is_production"`
}
type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
}

func NewClientConfig(configPath string) *ClientConfig {
	return loadConfig(configPath)
}

func loadConfig(configPath string) *ClientConfig {
	cfg := &ClientConfig{}
	data, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		panic(err)
	}
	return cfg
}
