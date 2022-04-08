package main

import (
	"portfolioManagement/server"
)

func main() {
	server := server.NewServer()
	server.Serve()
}
