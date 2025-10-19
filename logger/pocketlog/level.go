package pocketlog

// Level represents an available logging level.
type Level byte

const (
	// LevelDebug represents the lowest level of log, mostly used for
	// debugging purposes.
	LevelDebug Level = iota
	// LevelInfo represents a logging level that contains information
	// deemed valuable.
	LevelInfo
	// LevelWarn represents a logging level that contains information
	// deemed as warnings.
	LevelWarn
	// LevelError represents the highest logging level, only to be used
	// to trace errors.
	LevelError
	// LevelFatal represents the highest logging level, only to be used
	// to trace fatal errors.
	LevelFatal
)
