// Package log implements the backend's bespoke logging system.
//
// Backend log messages can be sent at four levels: Debug, Info, Warn, & Error. Messages can be formatted like printf.
// These log routines should be used instead of print statements. This package exists to allow easy implementation of
// more complex logging systems in the future, for instance if we wanted to save all log messages to a file, or send out
// an email if an error occurred.
package log

import (
	"fmt"
	"time"
)

const (
	levelDebug = iota
	levelInfo
	levelWarn
	levelError
)

// Debug logs a message at the debug level. Should be used for spamy log messages. Can be formatted like printf.
func Debug(format string, a ...any) { logAtLevel(levelDebug, format, a...) }

// Info logs a message at the info level. Should be used for important messages that do not indicate a problem. Can be
// formatted like printf.
func Info(format string, a ...any) { logAtLevel(levelInfo, format, a...) }

// Warn logs a message at the warn level. Should be used to indicate that something went wrong, but the backend is still
// operating as expected (e.g. malformed input). Can be formatted like printf.
func Warn(format string, a ...any) { logAtLevel(levelWarn, format, a...) }

// Error logs a message at the error level. Should be used to indicate that something unexpected happened that should
// probably be addressed. Can be formatted like printf.
func Error(format string, a ...any) { logAtLevel(levelError, format, a...) }

func logAtLevel(level int, format string, a ...any) {
	now := time.Now()
	levelStr := ""

	switch level {
	case levelDebug:
		levelStr = "[DEBUG]"
	case levelInfo:
		levelStr = "\u001B[34m[INFO]\u001B[0m "
	case levelWarn:
		levelStr = "\u001B[33m[WARN]\u001B[0m "
	case levelError:
		levelStr = "\u001B[31m[ERROR]\u001B[0m"
	default:
	}

	fmt.Printf("%s[%s] ", levelStr, now.Format(time.RFC3339))
	fmt.Printf(format, a...)
	fmt.Printf("\n")
}
