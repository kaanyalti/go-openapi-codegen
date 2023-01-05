package utils

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env            string
	DatabaseConfig DatabaseConfig
}

type DatabaseConfig struct {
	Database string `envconfig:"database"`
	User     string `envconfig:"database_user"`
	Password string `envconfig:"database_password"`
	Host     string `envconfig:"database_host"`
	Port     int    `envconfig:"database_port"`
}

func NewConfig() (Config, error) {
	var config Config

	err := envconfig.Process("", &config)
	if err != nil {
		log.Fatalln(err)
	}

	return config, nil
}
