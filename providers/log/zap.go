package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newZapLogger(conf *Config) (*Logger, error) {
	logger, err := zap.Config{
		Level:       zap.NewAtomicLevelAt(zapcore.Level(conf.Level)),
		Development: conf.Debug,
		Sampling:    &zap.SamplingConfig{Initial: 100, Thereafter: 100},
		Encoding:    "console",
		// Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zapcore.PanicLevel),
	)
	if err != nil {
		return nil, err
	}
	return &Logger{Logger: logger.Sugar()}, nil
}
