package golib

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger struct {
	driver *zap.Logger
}

func NewLogger(dir string) *Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "/var/log/" + dir + "/error.log",
		MaxSize:    25, // megabytes
		MaxBackups: 20,
		MaxAge:     28, // days
	})

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		w,
		zap.InfoLevel,
	)

	logger := new(Logger)
	logger.driver = zap.New(core, zap.AddCallerSkip(1), zap.AddCaller())
	return logger
}

func (logger Logger) Info() {

}

func (logger Logger) Error(msg string) {
	logger.driver.Error(msg)
}
