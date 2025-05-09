package logger

import (
	client_config "client_server/pkg/config"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ZapLogger struct {
	log    *zap.SugaredLogger
	config *client_config.ClientConfig
}

func initProductionConfig() zap.Config {
	cfg := zap.NewProductionConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.MessageKey = "message"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.DisableCaller = true
	return cfg
}

func initDevelopmentConfig() zap.Config {
	cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	cfg.EncoderConfig.TimeKey = "timestamp"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.MessageKey = "message"
	cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	cfg.DisableCaller = true
	return cfg
}

func NewZapLogger(config *client_config.ClientConfig) Logger {
	var cfg zap.Config
	if config.ServerConfig.IsProduction {
		cfg = initProductionConfig()
	} else {
		cfg = initDevelopmentConfig()
	}
	logger, _ := cfg.Build()
	sugar := logger.Sugar()
	return &ZapLogger{log: sugar, config: config}
}

func (l *ZapLogger) Debug(msg ...interface{}) {
	l.log.Debug(msg...)
}

func (l *ZapLogger) Info(msg ...interface{}) {
	l.log.Info(msg...)
}

func (l *ZapLogger) Warn(msg ...interface{}) {
	l.log.Warn(msg...)
}

func (l *ZapLogger) Error(msg ...interface{}) {
	l.log.Error(msg...)
}

func (l *ZapLogger) Fatal(msg ...interface{}) {
	l.log.Fatal(msg...)
}

func (l *ZapLogger) Flush() {
	_ = l.log.Sync()
}
