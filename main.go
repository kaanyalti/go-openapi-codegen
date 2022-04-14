package main

import (
	"portfolioManagement/infrastructure"
	"portfolioManagement/server"
	"portfolioManagement/ui/routers"
	"portfolioManagement/utils"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	config, err := utils.NewConfig()
	if err != nil {
		panic(err)
	}
	router := routers.NewRouter()
	logger := infrastructure.ProvideLogger()
	server := server.NewServer(config, router, logger)
	server.Serve()
}
