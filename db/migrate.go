package main

import (
	"flag"
	"fmt"
	"openapiCodegen/infrastructure"
	"openapiCodegen/utils"
	"strconv"

	"github.com/joho/godotenv"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	up := flag.Bool("up", false, "up migrate")
	down := flag.Bool("down", false, "down migrate")
	flag.Parse()

	if !*up && !*down {
		panic(fmt.Errorf("need to provide up or down flag for migrations"))
	}
	if *up && *down {
		panic(fmt.Errorf("cannot provide both up and down flag for migrations"))
	}

	godotenv.Load()
	config, err := utils.NewConfig()
	if err != nil {
		panic(err)
	}
	logger := infrastructure.ProvideLogger()
	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", config.DatabaseConfig.User, config.DatabaseConfig.Password, config.DatabaseConfig.Host, strconv.Itoa(config.DatabaseConfig.Port), config.DatabaseConfig.Database),
	)
	if err != nil {
		logger.Fatal(err)
	}

	if *up {
		if err := m.Up(); err != nil {
			logger.Fatal(err)
		}
	}

	if *down {
		if err := m.Down(); err != nil {
			logger.Fatal(err)
		}
	}
}
