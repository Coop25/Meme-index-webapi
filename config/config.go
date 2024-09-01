package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port        string `envconfig:"PORT" default:"8080"`
}

func LoadConfig() Config {
	var config Config
	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatalf("Failed to load config: %s", err)
	}
	return config
}
