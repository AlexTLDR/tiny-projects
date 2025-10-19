package pocketlog

import (
	"fmt"
	"io"
	"os"
)

// Logger is used to log information.
type Logger struct {
	threshold Level
	output    io.Writer
}

// Debugf formats and prints a message if the log level is debug or higher.
func (l *Logger) Debugf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}
	if l.threshold <= LevelDebug {
		l.logf("[DEBUG]", format, args...)
	}
}

func (l *Logger) Infof(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold <= LevelInfo {
		l.logf("[INFO]", format, args...)
	}
}

func (l *Logger) Warnf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold <= LevelWarn {
		l.logf("[WARN]", format, args...)
	}
}

func (l *Logger) Errorf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold <= LevelError {
		l.logf("[ERROR]", format, args...)
	}
}

func (l *Logger) Fatalf(format string, args ...any) {
	if l.output == nil {
		l.output = os.Stdout
	}

	if l.threshold > LevelFatal {
		l.logf("[FATAL]", format, args...)
	}
}

func (l *Logger) logf(level string, format string, args ...any) {
	_, _ = fmt.Fprintf(l.output, level+": "+format+"\n", args...)
}

// New returns you a logger, ready to log at the required threshold.
// The default output is Stdout.
func New(threshold Level, output io.Writer) *Logger {
	return &Logger{
		threshold: threshold,
		output:    output,
	}
}
