package main

import (
	"portfolioManagement/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	server := server.NewServer()
	server.Serve()
}
