package server

import (
	"portfolioManagement/infrastructure"
	"portfolioManagement/utils"
)

func ProvideLogger() utils.Logger {
	return infrastructure.NewLocalLogger()
}
