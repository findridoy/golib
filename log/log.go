package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var infoLogger *zap.Logger
var errorLogger *zap.Logger

func Init(dir string) {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   "/var/log/" + dir + "/info.log",
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

	infoLogger = zap.New(core)

	w = zapcore.AddSync(&lumberjack.Logger{
		Filename:   "/var/log/" + dir + "/error.log",
		MaxSize:    5, // megabytes
		MaxBackups: 10,
		MaxAge:     28, // days
	})

	encoderConfig = zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core = zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		w,
		zap.InfoLevel,
	)

	errorLogger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func Info(msg string) {
	infoLogger.Info(msg)
}

func Error(msg string) {
	errorLogger.Error(msg)
}

func Warn(msg string) {
	errorLogger.Warn(msg)
}
