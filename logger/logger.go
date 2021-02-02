package logger

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	envLogLevel  = "LOG_LEVEL"
	envLogOutput = "LOG_OUTPUT"
)

var (
	log logger
)

type loggerInterface interface {
	Print(...interface{})
	Printf(string, ...interface{})
	ErrorPrint(error, ...interface{})
	ErrorPrintf(string, error, ...interface{})
}

type logger struct {
	log *zap.Logger
}

func init() {
	//config the logger
	logConfig := zap.Config{
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:     "level",
			TimeKey:      "time",
			MessageKey:   "msg",
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	// Build the log engine
	var err error
	if log.log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

func Logger() loggerInterface {
	return log
}

func getLevel() zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(os.Getenv(envLogLevel))) {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "error":
		return zap.ErrorLevel
	default:
		return zap.InfoLevel
	}
}

func getOutput() string {
	output := strings.TrimSpace(os.Getenv(envLogOutput))
	if output == "" {
		return "stdout"
	}
	return output
}

func (l logger) Print(v ...interface{}) {
	Info(fmt.Sprintf("%v", v))
}

func (l logger) Printf(format string, v ...interface{}) {
	if len(v) == 0 {
		Info(format)
		return
	}

	Info(fmt.Sprintf(format, v...))
}

func (l logger) ErrorPrint(err error, v ...interface{}) {
	Error(fmt.Sprintf("%v", v), err)
}

func (l logger) ErrorPrintf(format string, err error, v ...interface{}) {
	if len(v) == 0 {
		Error(format, err)
		return
	}

	Error(fmt.Sprintf("%v", v), err)
}

// Info is a log function which log an custom info type
func Info(msg string, tags ...zap.Field) {
	log.log.Info(msg, tags...)
	log.log.Sync()
}

// Error is a log function which log an custom error type
func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))
	log.log.Error(msg, tags...)
	log.log.Sync()
}
