package infrastructure

import (
	"log"
	"portfolioManagement/utils"
)

type localLogger struct{}

func NewLocalLogger() utils.Logger {
	return &localLogger{}
}

func (l *localLogger) Info(v ...interface{}) {
	log.Println(v)
}

func (l *localLogger) Error(v ...interface{}) {
	log.Println(v)
}

func (l *localLogger) Fatal(v ...interface{}) {
	log.Fatalln(v)
}

func (l *localLogger) Debug(v ...interface{}) {
	log.Println(v)
}
