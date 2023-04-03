package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var logger *zap.Logger

func Init() {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder

	logLevel := zap.InfoLevel
	if os.Getenv("log_level") == "debug" {
		logLevel = zap.DebugLevel
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(os.Stdout),
		logLevel,
	)

	logger = zap.New(core)
}

func Debug(message string, field ...zapcore.Field) {
	logger.Debug(message, field...)
}
func Info(message string, field ...zapcore.Field) {
	logger.Info(message, field...)
}
func Warn(message string, field ...zapcore.Field) {
	logger.Warn(message, field...)
}
func Error(message string, field ...zapcore.Field) {
	logger.Error(message, field...)
}
func Panic(message string, field ...zapcore.Field) {
	logger.Panic(message, field...)
}
