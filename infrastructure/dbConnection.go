package infrastructure

import (
	"portfolioManagement/utils"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDbConnection(logger utils.Logger) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	retries := 5
	isConnected := false

	for !isConnected {
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN: "host=localhost user=root password=root dbname=portfoliomanagement port=5432 sslmode=disable",
		}), &gorm.Config{})

		retries--
		if retries <= 0 {
			logger.Fatal(err)
		}

		if err != nil {
			logger.Info("Couldn't connect to db, retrying")
			time.Sleep(time.Second * 3)
		}
	}

	return db, nil
}
