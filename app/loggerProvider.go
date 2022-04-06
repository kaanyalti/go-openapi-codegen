package app

import (
	"portfolioManagement/infrastructure"
	utils "portfolioManagement/utils/logger"
)

func ProvideLogger() utils.Logger {
	return infrastructure.NewLocalLogger()
}
