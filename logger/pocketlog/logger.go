package pocketlog

import (
	"fmt"
)

type Logger struct {
	threshold Level
}

// Debugf formats and prints a message if the log level is debug or higher
func (l *Logger) Debugf(format string, args ...any) {
	if l.threshold > LevelDebug {
		return
	}

	_, _ = fmt.Printf(format+"\n", args...)
}

func (l *Logger) Infof(format string, args ...any) {
	if l.threshold > LevelInfo {
		return
	}
	_, _ = fmt.Printf(format+"\n", args...)
}

func (l *Logger) Warnf(format string, args ...any) {
	if l.threshold > LevelWarn {
		return
	}
	_, _ = fmt.Printf(format+"\n", args...)
}

func (l *Logger) Errorf(format string, args ...any) {
	if l.threshold > LevelError {
		return
	}
	_, _ = fmt.Printf(format+"\n", args...)
}

func (l *Logger) Fatalf(format string, args ...any) {
	if l.threshold > LevelFatal {
		return
	}
	_, _ = fmt.Printf(format+"\n", args...)
}

func New(threshold Level) *Logger {
	return &Logger{threshold: threshold}
}
