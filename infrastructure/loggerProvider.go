package infrastructure

import (
	"portfolioManagement/utils"
)

func ProvideLogger() utils.Logger {
	return NewLocalLogger()
}
