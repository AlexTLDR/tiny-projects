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

// logf prints the message to the output.
// Add decorations here, if any.
func (l *Logger) logf(format string, args ...any) {
	_, _ = fmt.Fprintf(l.output, format+"\n", args...)
}
// Logf logs a message at the specified level.
func (l *Logger) Logf(lvl Level, format string, args ...any){
	if l.threshold > lvl {
		return
	}
	l.logf(format, args...)
}

// New returns you a logger, ready to log at the required threshold.
// The default output is Stdout.
func New(threshold Level, opts ...Option) *Logger {
	lgr := &Logger{threshold: threshold, output: os.Stdout}

	for _, configFunc := range opts {
		configFunc(lgr)
	}
	return lgr
}
