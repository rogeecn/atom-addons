package log

import (
	"strings"
)

type LevelWriter struct {
	Level Level
}

func (w LevelWriter) Write(p []byte) (n int, err error) {
	str := strings.TrimSpace(string(p))
	switch w.Level {
	case InfoLevel:
		Info(str)
	case ErrorLevel:
		Error(str)
	}
	return len(p), nil
}
