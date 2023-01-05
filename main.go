package main

import (
	"openapiCodegen/infrastructure"
	"openapiCodegen/server"
	"openapiCodegen/ui/routers"
	"openapiCodegen/utils"

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
