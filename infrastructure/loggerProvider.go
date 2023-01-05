package infrastructure

import (
	"openapiCodegen/utils"
)

func ProvideLogger() utils.Logger {
	return NewLocalLogger()
}
