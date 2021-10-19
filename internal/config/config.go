package config

import (
	"github.com/kelseyhightower/envconfig"
)

// SERVICENAME is an environment variables prefix
const SERVICENAME = "BookStore"

// AppConfig contains all the config variables for the application
type AppConfig struct {
	DbConfig  DatabaseConfig
	LogConfig LogConfig
}

// DatabaseConfig contains variables, that are required for a database connection
type DatabaseConfig struct {
	Dialect          string `split_words:"true" required:"true"`
	ConnectionString string `split_words:"true" required:"true"`
}

// LogConfig contains variables, that define log behavior
type LogConfig struct {
	WriteToFile bool   `split_words:"true"`
	Filepath    string `split_words:"true"`
	Level       string `split_words:"true"`
}

// Load sets environment variables into AppConfig structure
func (c *AppConfig) Load(serviceName string) error {
	return envconfig.Process(serviceName, c)
}
