package log

import (
	"errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var _logger *logger
var errlogger *errLogger

func NewConfig(dir string) Configure {
	return &config{
		Dir:  dir,
		Path: ".",
	}
}

func Init(c Configure) error {
	config, ok := c.(*config)
	if !ok {
		return errors.New("invalid config type")
	}
	lumberjackLogger := &lumberjack.Logger{
		Filename:   config.Path + "/" + config.Dir + "/process.log",
		MaxSize:    25, // megabytes
		MaxBackups: 20,
		MaxAge:     28, // days
	}

	_, err := lumberjackLogger.Write([]byte(""))
	if err != nil {
		return err
	}

	w := zapcore.AddSync(lumberjackLogger)

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		w,
		zap.InfoLevel,
	)

	logger := new(logger)
	logger.driver = zap.New(core)
	_logger = logger

	lumberjackLogger = &lumberjack.Logger{
		Filename:   config.Path + "/" + config.Dir + "/error.log",
		MaxSize:    25, // megabytes
		MaxBackups: 20,
		MaxAge:     28, // days
	}
	_, err = lumberjackLogger.Write([]byte(""))
	if err != nil {
		return err
	}

	w = zapcore.AddSync(lumberjackLogger)

	encoderConfig = zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	core = zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		w,
		zap.ErrorLevel,
	)

	eLogger := new(errLogger)
	eLogger.driver = zap.New(core)
	errlogger = eLogger

	return nil
}

type Configure interface {
	SetPath(path string)
}
type config struct {
	Dir  string
	Path string
}

func (c *config) SetPath(path string) {
	c.Path = path
}

type logger struct {
	driver *zap.Logger
}

func Info(msg string) {
	_logger.driver.Info(msg)
}

type errLogger struct {
	driver *zap.Logger
}

func Error(msg string) {
	errlogger.driver.Error(msg)
}
