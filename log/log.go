package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.Logger

func InitLog(appName string, level zapcore.Level) {
	InitLogWithPath("logs/", appName, level)
}

func InitLogWithPath(path, appName string, level zapcore.Level) {
	filePath := path + appName + ".log"
	jLoger := &lumberjack.Logger{
		Filename: filePath,
		MaxSize:  100, //MB
	}
	w := zapcore.AddSync(jLoger)

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config),
		w,
		level,
	)

	log = zap.New(core)
	log = log.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

func Debugf(template string, args ...interface{}) {
	log.Debug(fmt.Sprintf(template, args...))
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

func Infof(template string, args ...interface{}) {
	log.Info(fmt.Sprintf(template, args...))
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(msg string, fields ...zap.Field) {
	log.Warn(msg, fields...)
}

func Warnf(template string, args ...interface{}) {
	log.Warn(fmt.Sprintf(template, args...))
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

func Errorf(template string, args ...interface{}) {
	log.Error(fmt.Sprintf(template, args...))
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func Fatal(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

func Fatalf(template string, args ...interface{}) {
	log.Fatal(fmt.Sprintf(template, args...))
}
