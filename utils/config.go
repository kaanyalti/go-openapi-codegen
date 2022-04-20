package utils

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Env            string
	DatabaseConfig *DatabaseConfig
}

type DatabaseConfig struct {
	Database string
	User     string
	Password string
	Host     string
	Port     int
}

func NewConfig() (*Config, error) {
	var missingEnvs []string
	var err error
	dbConfig := &DatabaseConfig{}
	config := &Config{
		DatabaseConfig: dbConfig,
	}

	env, ok := os.LookupEnv("ENV")
	if !ok {
		missingEnvs = append(missingEnvs, "ENV")
	}

	dbName, ok := os.LookupEnv("DB_NAME")
	if !ok {
		missingEnvs = append(missingEnvs, "DB_NAME")
	}
	dbUser, ok := os.LookupEnv("DB_USER")
	if !ok {
		missingEnvs = append(missingEnvs, "DB_USER")
	}
	dbPassword, ok := os.LookupEnv("DB_PASSWORD")
	if !ok {
		missingEnvs = append(missingEnvs, "DB_PASSWORD")
	}
	dbHost, ok := os.LookupEnv("DB_HOST")
	if !ok {
		missingEnvs = append(missingEnvs, "DB_HOST")
	}
	dbPort, ok := os.LookupEnv("DB_PORT")
	if !ok {
		missingEnvs = append(missingEnvs, "DB_PORT")
	}

	if len(missingEnvs) != 0 {
		var errMessage strings.Builder
		for _, envVar := range missingEnvs {
			errMessage.WriteString(fmt.Sprintf("%s\n", envVar))
		}
		return nil, fmt.Errorf(fmt.Sprintf("The following environment variables are not provided:\n%s", errMessage.String()))
	}

	config.Env = env

	dbConfig.Database = dbName
	dbConfig.User = dbUser
	dbConfig.Password = dbPassword
	dbConfig.Host = dbHost
	dbConfig.Port, err = strconv.Atoi(dbPort)
	if err != nil {
		return nil, err
	}

	return config, nil
}
