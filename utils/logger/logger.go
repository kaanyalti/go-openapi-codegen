package utils

type Logger interface {
	Info(...interface{})
	Error(...interface{})
	Debug(...interface{})
}
