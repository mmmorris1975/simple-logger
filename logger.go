package simple_logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const (
	// NONE is the log level to indicate no log output should be generated
	NONE = iota
	// FATAL log level will only output messages at the FATAL level, this is the highest log level
	FATAL
	// ERROR log level outputs messages at the ERROR level or higher
	ERROR
	// WARN log level outputs messages at the WARN level or higher
	WARN
	// INFO log level outputs messages at the INFO level or higher
	INFO
	// DEBUG log level outputs messages at the DEBUG level or higher, this is the lowest log level
	DEBUG
)

var levels = []string{"NONE", "FATAL", "ERROR", "WARN", "INFO", "DEBUG"}

// Logger is a logging object which provides leveled logging using the stdlib log.Logger
type Logger struct {
	Level uint
	*log.Logger
}

// StdLogger is a shortcut to get a logger which logs to stderr with the stdlib standard logging flags (log.LstdFlags)
var StdLogger = NewLogger(os.Stderr, "", log.LstdFlags)

// NewLogger provides a way to customize a logger object by specifying the output io.Writer, a desired prefix (empty string
// for no prefix, and any flags which will control the output decorations (see the constants in the stdlib log package).
// The returned logger is set to the INFO level by default, but can be modified by updating the Level field, or calling
// ParseLevel()/SetLevel().
func NewLogger(out io.Writer, prefix string, flag int) *Logger {
	l := &Logger{Level: INFO, Logger: log.New(out, prefix, flag)}
	return l
}

// ParseLevel accepts a string as a log level name, and returns the corresponding int value for the level
func ParseLevel(level string) (uint, error) {
	for i, v := range levels {
		if strings.EqualFold(v, level) {
			return uint(i), nil
		}
	}

	return 0, fmt.Errorf("invalid log Level '%s'", level)
}

// SetLevel will set the logger to only output messages at the provided level or higher
func (l *Logger) SetLevel(level uint) {
	l.Level = level
}

// Fatalf logs a formatted message string at the FATAL level, and exits (via os.Exit)
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.writeLogf(FATAL, format, v...)
	os.Exit(1)
}

// Fatal logs a message string at the FATAL level, and exits (via os.Exit)
func (l *Logger) Fatal(v ...interface{}) {
	l.writeLogln(FATAL, v...)
	os.Exit(1)
}

// Fatalln logs a message string at the FATAL level, and exits (via os.Exit)
func (l *Logger) Fatalln(v ...interface{}) {
	l.writeLogln(FATAL, v...)
	os.Exit(1)
}

// Errorf logs a formatted message string at the ERROR level
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.writeLogf(ERROR, format, v...)
}

// Error logs a message string at the ERROR level
func (l *Logger) Error(v ...interface{}) {
	l.writeLogln(ERROR, v...)
}

// Errorln logs a message string at the ERROR level
func (l *Logger) Errorln(v ...interface{}) {
	l.writeLogln(ERROR, v...)
}

// Warnf logs a formatted message string at the WARN level
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.writeLogf(WARN, format, v...)
}

// Warn logs a message string at the WARN level
func (l *Logger) Warn(v ...interface{}) {
	l.writeLogln(WARN, v...)
}

// Warnln logs a message string at the WARN level
func (l *Logger) Warnln(v ...interface{}) {
	l.writeLogln(WARN, v...)
}

// Infof logs a formatted message string at the INFO level
func (l *Logger) Infof(format string, v ...interface{}) {
	l.writeLogf(INFO, format, v...)
}

// Info logs a message string at the INFO level
func (l *Logger) Info(v ...interface{}) {
	l.writeLogln(INFO, v...)
}

// Infoln logs a message string at the INFO level
func (l *Logger) Infoln(v ...interface{}) {
	l.writeLogln(INFO, v...)
}

// Debugf logs a formatted message string at the DEBUG level
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.writeLogf(DEBUG, format, v...)
}

// Debug logs a message string at the DEBUG level
func (l *Logger) Debug(v ...interface{}) {
	l.writeLogln(DEBUG, v...)
}

// Debugln logs a message string at the DEBUG level
func (l *Logger) Debugln(v ...interface{}) {
	l.writeLogln(DEBUG, v...)
}

// Panicf outputs a formatted message string, and calls panic(), bypassing log level checking
func (l *Logger) Panicf(format string, v ...interface{}) {
	// Output directly for all Panic*() calls, avoid Level checking
	msg := fmt.Sprintf(format, v...)
	l.Output(3, fmt.Sprintf("PANIC %s", msg))
	panic(msg)
}

// Panic outputs the message, and calls panic(), bypassing log level checking
func (l *Logger) Panic(v ...interface{}) {
	msg := fmt.Sprint(v...)
	l.Output(3, fmt.Sprintf("PANIC %s", msg))
	panic(msg)
}

// Panicln outputs the message, and calls panic(), bypassing log level checking
func (l *Logger) Panicln(v ...interface{}) {
	msg := fmt.Sprintln(v...)
	l.Output(3, fmt.Sprintf("PANIC %s", msg))
	panic(msg)
}

// Printf outputs a formatted message string, bypassing log level checking
func (l *Logger) Printf(format string, v ...interface{}) {
	// Print*() logs an "un-leveled" message
	l.Output(3, fmt.Sprintf(format, v...))
}

// Print outputs the message, bypassing log level checking
func (l *Logger) Print(v ...interface{}) {
	l.Output(3, fmt.Sprintln(v...))
}

// Println outputs the message, bypassing log level checking
func (l *Logger) Println(v ...interface{}) {
	l.Output(3, fmt.Sprintln(v...))
}

// Logf outputs a formatted message string, at the configured log level.  Used for compatibility with other logging interfaces.
// Will require wrapping the call to this method in a conditional if you wish to control what is output
//
// Example:
//   l.SetLevel(DEBUG)
//   l.Logf("%s", "message")
//
// will write "DEBUG message" out, and if the level was set to "WARN" it would write "WARN message".
// To control the output on the caller side (debugging), it would be necessary to do something similar to:
//
//  l.SetLevel(ERROR)
//  if l.Level >= DEBUG {
//    l.Logf("%s", "message")
//  }
//
// so that the l.Logf() call only fires if the logging level is at least DEBUG, but any other messages are written at the
// "ERROR" level
func (l *Logger) Logf(format string, v ...interface{}) {
	l.writeLogf(l.Level, format, v...)
}

// Log outputs the message, at the configured log level.  Used for compatibility with other logging interfaces.
// See documentation for Logf() about controlling log output on the caller side
func (l *Logger) Log(v ...interface{}) {
	l.writeLogln(l.Level, v...)
}

func (l *Logger) writeLogf(level uint, format string, v ...interface{}) error {
	if l.Level >= level {
		q := []interface{}{levels[level]}
		q = append(q, v...)
		return l.Output(3, fmt.Sprintf("%s "+format, q...))
	}
	return nil
}

func (l *Logger) writeLogln(level uint, v ...interface{}) error {
	if l.Level >= level {
		q := []interface{}{levels[level]}
		q = append(q, v...)
		return l.Output(3, fmt.Sprintln(q...))
	}
	return nil
}
