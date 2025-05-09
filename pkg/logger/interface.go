package logger

type Logger interface {
	Info(...interface{})
	Error(...interface{})
	Debug(...interface{})
	Warn(...interface{})
	Fatal(...interface{})
	Flush()
}
