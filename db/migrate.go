package db

import (
	"portfolioManagement/infrastructure"
	"portfolioManagement/utils"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config, err := utils.NewConfig()
	if err != nil {
		panic(err)
	}
	logger := infrastructure.ProvideLogger()
}
