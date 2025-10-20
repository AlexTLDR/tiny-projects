package pocketlog_test

import (
	"os"

	"github.com/AlexTLDR/tiny-projects/logger/pocketlog"
)

func ExampleLogger_debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug, os.Stdout)
	debugLogger.Debugf("Hello, %s", "world")
}
