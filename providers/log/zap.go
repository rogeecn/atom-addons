package log

import (
	"go.uber.org/zap"
)

func newZapLogger(conf *Config) (*Logger, error) {
	logger, err := zap.NewDevelopment()
	if err != nil {
		return nil, err
	}
	return &Logger{logger: logger.Sugar()}, nil
}
