package common

import (
	"fmt"
	"os"

	logger "github.com/soulteary/logger-kit"
)

// LogLevel represents the logging level (owlmail legacy: silent / normal / verbose).
type LogLevel int

const (
	// LogLevelSilent suppresses all logs except errors
	LogLevelSilent LogLevel = iota
	// LogLevelNormal shows normal logs
	LogLevelNormal
	// LogLevelVerbose shows detailed logs
	LogLevelVerbose
)

// owlmailToLoggerLevel maps owlmail LogLevel to logger-kit Level.
func owlmailToLoggerLevel(l LogLevel) logger.Level {
	switch l {
	case LogLevelSilent:
		return logger.ErrorLevel // only errors in silent mode
	case LogLevelNormal:
		return logger.InfoLevel
	case LogLevelVerbose:
		return logger.DebugLevel
	default:
		return logger.InfoLevel
	}
}

// InitLogger initializes the global logger with the specified level using logger-kit.
func InitLogger(level LogLevel) {
	kitLevel := owlmailToLoggerLevel(level)
	l := logger.New(logger.Config{
		Level:       kitLevel,
		Output:      os.Stdout,
		Format:      logger.FormatConsole,
		ServiceName: "owlmail",
	})
	logger.SetDefault(l)
}

// Log prints a log message at info level (normal and verbose).
func Log(format string, v ...interface{}) {
	logger.Default().Info().Msg(fmt.Sprintf(format, v...))
}

// Verbose prints a log message at debug level (verbose only).
func Verbose(format string, v ...interface{}) {
	logger.Default().Debug().Msg(fmt.Sprintf(format, v...))
}

// Error prints an error message (always shown).
func Error(format string, v ...interface{}) {
	logger.Default().Error().Msg(fmt.Sprintf(format, v...))
}

// Fatal logs a fatal error via the error handler (returns error in tests, exits in production).
// Logs at Error level first, then delegates to the error handler.
func Fatal(format string, v ...interface{}) error {
	logger.Default().Error().Msg(fmt.Sprintf(format, v...))
	return GetErrorHandler().Fatal(format, v...)
}
