package pocketlog

type Logger struct {
	threshold Level
}

func (l *Logger) Debugf(format string, args ...any) {

}

func (l *Logger) Infof(format string, args ...any) {

}

func (l *Logger) Warnf(format string, args ...any) {

}

func (l *Logger) Errorf(format string, args ...any) {

}

func (l *Logger) Fatalf(format string, args ...any) {

}

func New(threshold Level) *Logger {
	return &Logger{threshold: threshold}
}
