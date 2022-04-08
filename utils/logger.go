package utils

type Logger interface {
	Info(...interface{})
	Error(...interface{})
	Fatal(...interface{})
	Debug(...interface{})
}
