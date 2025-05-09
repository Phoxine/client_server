package logger

import (
	client_config "client_server/pkg/config"

	"github.com/sirupsen/logrus"
)

type LogrusLogger struct {
	log    *logrus.Logger
	config *client_config.ClientConfig
}

func NewLogrusLogger(config *client_config.ClientConfig) Logger {
	logger := &LogrusLogger{log: logrus.New(), config: config}
	logger.setLogger()
	return logger
}

func (l *LogrusLogger) setLogger() {
	if l.config.ServerConfig.IsProduction {
		l.log.SetFormatter(&logrus.JSONFormatter{})
		l.log.SetLevel(logrus.InfoLevel)
	} else {
		l.log.SetFormatter(&logrus.TextFormatter{})
		l.log.SetLevel(logrus.DebugLevel)
	}
}

// func transferLogLevel(logLevel string) logrus.Level {
// 	switch logLevel {
// 	case "debug":
// 		return logrus.DebugLevel
// 	case "info":
// 		return logrus.InfoLevel
// 	case "warn":
// 		return logrus.WarnLevel
// 	case "error":
// 		return logrus.ErrorLevel
// 	case "fatal":
// 		return logrus.FatalLevel
// 	default:
// 		return logrus.DebugLevel
// 	}
// }

func (l *LogrusLogger) Info(msg ...interface{}) {
	l.log.Info(msg...)
}

func (l *LogrusLogger) Error(msg ...interface{}) {
	l.log.Error(msg...)
}

func (l *LogrusLogger) Debug(msg ...interface{}) {
	l.log.Debug(msg...)
}

func (l *LogrusLogger) Warn(msg ...interface{}) {
	l.log.Warn(msg...)
}

func (l *LogrusLogger) Fatal(msg ...interface{}) {
	l.log.Fatal(msg...)
}

func (l *LogrusLogger) Flush() {
}
