package utils

import "os"

type Config struct {
	Env string
}

func NewConfig(logger Logger) *Config {
	var config *Config

	env, ok := os.LookupEnv("ENV")

	return config
}
