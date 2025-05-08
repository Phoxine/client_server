package logger

type Logger interface {
	Info(string)
	Error(string)
	Debug(string)
	Warn(string)
	Fatal(string)
}
