package pocketlog_test

import (
	"testing"

	"github.com/AlexTLDR/tiny-projects/logger/pocketlog"
)

const (
	debugMessage   = "Debug message test"
	infoMessage    = "Info message test"
	warningMessage = "Warning message test"
	errorMessage   = "Error message test"
	fatalMessage   = "Fatal message test"
)

// testWriter is a struct that implements io.Writer.
// We use it to validate that we can write to a specific output.
type testWriter struct {
	contents string
}

func (t *testWriter) Write(p []byte) (n int, err error) {
	t.contents += string(p)
	return len(p), nil
}

func ExampleLogger_Logf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Logf(pocketlog.LevelDebug, "Hello, %s", "world")
}

func TestLogger_DebugfInfofWarningfErrorfFatalf(t *testing.T) {
	type testCase struct {
		level    pocketlog.Level
		expected string
	}

	tt := map[string]testCase{
		"debug": {
			level:    pocketlog.LevelDebug,
			expected: debugMessage + "\n" + infoMessage + "\n" + warningMessage + "\n" + errorMessage + "\n" + fatalMessage + "\n",
		},
		"info": {
			level:    pocketlog.LevelInfo,
			expected: infoMessage + "\n" + warningMessage + "\n" + errorMessage + "\n" + fatalMessage + "\n",
		},
		"warning": {
			level:    pocketlog.LevelWarn,
			expected: warningMessage + "\n" + errorMessage + "\n" + fatalMessage + "\n",
		},
		"error": {
			level:    pocketlog.LevelError,
			expected: errorMessage + "\n" + fatalMessage + "\n",
		},
		"fatal": {
			level:    pocketlog.LevelFatal,
			expected: fatalMessage + "\n",
		},
	}

	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			tw := &testWriter{}

			testedLogger := pocketlog.New(tc.level, pocketlog.WithOutput(tw))
			testedLogger.Logf(pocketlog.LevelDebug, debugMessage)
			testedLogger.Logf(pocketlog.LevelInfo, infoMessage)
			testedLogger.Logf(pocketlog.LevelWarn, warningMessage)
			testedLogger.Logf(pocketlog.LevelError, errorMessage)
			testedLogger.Logf(pocketlog.LevelFatal, fatalMessage)

			if tw.contents != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, tw.contents)
			}
		})
	}
}
