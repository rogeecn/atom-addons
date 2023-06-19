package log

import (
	"github.com/rogeecn/atom/container"
	"github.com/rogeecn/atom/utils/opt"

	"go.uber.org/zap"
)

func Provide(opts ...opt.Option) error {
	o := opt.New(opts...)
	var config Config
	if err := o.UnmarshalConfig(&config); err != nil {
		return err
	}

	logger, err := newZapLogger(&config)
	if err != nil {
		return err
	}
	defaultLogger = logger
	return container.Container.Provide(func() (*Logger, error) {
		return defaultLogger, nil
	}, o.DiOptions()...)
}

var defaultLogger *Logger

type Logger struct {
	Logger *zap.SugaredLogger
}

func (l *Logger) LevelWriter(level Level) *LevelWriter {
	return &LevelWriter{Level: level}
}

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	defaultLogger.Logger.Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	defaultLogger.Logger.Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	defaultLogger.Logger.Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	defaultLogger.Logger.Error(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(args ...interface{}) {
	defaultLogger.Logger.DPanic(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	defaultLogger.Logger.Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	defaultLogger.Logger.Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	defaultLogger.Logger.Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	defaultLogger.Logger.Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	defaultLogger.Logger.Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	defaultLogger.Logger.Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
	defaultLogger.Logger.DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	defaultLogger.Logger.Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	defaultLogger.Logger.Fatalf(template, args...)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//
//	s.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
	defaultLogger.Logger.Debugw(msg, keysAndValues...)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	defaultLogger.Logger.Infow(msg, keysAndValues...)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	defaultLogger.Logger.Warnw(msg, keysAndValues...)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	defaultLogger.Logger.Errorw(msg, keysAndValues...)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...interface{}) {
	defaultLogger.Logger.DPanicw(msg, keysAndValues...)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
	defaultLogger.Logger.Panicw(msg, keysAndValues...)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	defaultLogger.Logger.Fatalw(msg, keysAndValues...)
}

// Sync flushes any buffered log entries.
func Sync() error {
	return defaultLogger.Logger.Sync()
}
