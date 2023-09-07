package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger

func Init(level string) {
	logger = logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	parsedLevel, err := logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}

	logger.SetLevel(parsedLevel)
}

func Panic(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

func Fatal(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

func Error(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Warn(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Info(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Debug(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Trace(format string, args ...interface{}) {
	logger.Tracef(format, args...)
}
