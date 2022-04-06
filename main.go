package main

import (
	"portfolioManagement/infrastructure"
	utils "portfolioManagement/utils/logger"
)

var Logger utils.Logger

func main() {
	Logger = infrastructure.NewLocalLogger()
}
