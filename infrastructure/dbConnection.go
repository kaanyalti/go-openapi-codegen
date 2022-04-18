package infrastructure

import (
	"fmt"
	"portfolioManagement/utils"
	"strconv"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDbConnection(config *utils.Config, logger utils.Logger) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	retries := 5
	isConnected := false

	for !isConnected {
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN: fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", config.DatabaseConfig.Host, config.DatabaseConfig.User, config.DatabaseConfig.Password, config.DatabaseConfig.Database, strconv.Itoa(config.DatabaseConfig.Port)),
		}), &gorm.Config{})

		retries--
		if retries <= 0 {
			logger.Fatal(err)
		}

		if err != nil {
			logger.Info("Couldn't connect to db, retrying")
			time.Sleep(time.Second * 3)
		} else {
			logger.Info("Successfully connected to db")
			isConnected = true
		}
	}

	return db, nil
}
