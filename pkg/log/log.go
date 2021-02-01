package log

import (
	"go.uber.org/zap"
)

func init() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	zap.ReplaceGlobals(logger)
}

func Error(args ...interface{}) {
	zap.S().Error(args)
}

func Errorf(template string, args ...interface{}) {
	zap.S().Errorf(template, args...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	zap.S().Errorw(msg, keysAndValues...)
}

func Info(args ...interface{}) {
	zap.S().Info(args...)
}

func Infof(template string, args ...interface{}) {
	zap.S().Infof(template, args...)
}

func Fatal(args ...interface{}) {
	zap.S().Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	zap.S().Fatalf(template, args...)
}
