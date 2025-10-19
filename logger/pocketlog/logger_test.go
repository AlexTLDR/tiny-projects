package pocketlog_test

import "github.com/AlexTLDR/tiny-projects/logger/pocketlog"

func ExampleLogger_debugf() {
	debugLogger := pocketlog.New(pocketlog.LevelDebug)
	debugLogger.Debugf("Hello, %s", "world")
}
