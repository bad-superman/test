package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var sugarLogger *zap.SugaredLogger

func init() {
	InitLogger()
}

func InitLogger() {

	encoder := getEncoder()
	writeSyncer := getLogWriter()
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.InfoLevel)

	// zap.AddCaller()  添加将调用函数信息记录到日志中的功能。
	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // 修改时间编码器

	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// NewConsoleEncoder 打印更符合人们观察的方式
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &Logger{
		Filename:   "./test.log",
		MaxSize:    1024,
		MaxBackups: 200,
		MaxAge:     7,
		Compress:   false,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	sugarLogger.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	sugarLogger.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	sugarLogger.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	sugarLogger.Error(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panicsugarLogger. (See DPanicLevel for detailsugarLogger.)
func DPanic(args ...interface{}) {
	sugarLogger.DPanic(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panicsugarLogger.
func Panic(args ...interface{}) {
	sugarLogger.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls osugarLogger.Exit.
func Fatal(args ...interface{}) {
	sugarLogger.Error(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	sugarLogger.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	sugarLogger.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	sugarLogger.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	sugarLogger.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panicsugarLogger. (See DPanicLevel for detailsugarLogger.)
func DPanicf(template string, args ...interface{}) {
	sugarLogger.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panicsugarLogger.
func Panicf(template string, args ...interface{}) {
	sugarLogger.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls osugarLogger.Exit.
func Fatalf(template string, args ...interface{}) {
	sugarLogger.Fatalf(template, args...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//  sugarLogger.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
	sugarLogger.Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	sugarLogger.Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	sugarLogger.Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	sugarLogger.Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panicsugarLogger. (See DPanicLevel for detailsugarLogger.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...interface{}) {
	sugarLogger.DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panicsugarLogger. The
// variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
	sugarLogger.Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls osugarLogger.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	sugarLogger.Fatalw(msg, keysAndValues...)
}

// Sync flushes any buffered log entriesugarLogger.
func Sync() error {
	return sugarLogger.Sync()
}
