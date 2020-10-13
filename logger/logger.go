package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type logger struct {
	*log.Logger
	level uint8
	test  bool
}

// StdLogger is a shortcut to get a logger which logs to stderr with the stdlib standard logging flags (log.LstdFlags)
var StdLogger = NewLogger(os.Stderr, "", log.LstdFlags)

// NewLogger returns a new logger object which will write output to the provided io.Writer.  The 'out', 'prefix' and
// 'flag' arguments are exactly the same as the arguments for the New() method in the golang log package.  This logger
// inherits from the golang log.Logger, so all methods of that type are available as well, including Print* and Panic*.
//
// The logger is initially created with the level set to INFO, but can be modified by calling ParseLevel()/SetLevel().
func NewLogger(out io.Writer, prefix string, flag int) *logger {
	return &logger{level: INFO, Logger: log.New(out, prefix, flag)}
}

// ParseLevel accepts a string as a log level name, and returns the corresponding int value for the level
func ParseLevel(level string) (uint8, error) {
	for i, v := range levels {
		if strings.EqualFold(v, level) {
			return uint8(i), nil
		}
	}

	return 0, fmt.Errorf("invalid log Level '%s'", level)
}

// SetLevel will set the logger to only output messages at the provided level or higher
func (l *logger) SetLevel(level uint8) {
	l.level = level
}

// WithLevel is a fluent method for setting the level for a logger
func (l *logger) WithLevel(level uint8) *logger {
	l.level = level
	return l
}

// Debug logs a message string at the DEBUG level
func (l *logger) Debug(v ...interface{}) {
	l.writeLog(DEBUG, fmt.Sprint(v...))
}

// Debugf logs a formatted message string at the DEBUG level
func (l *logger) Debugf(format string, v ...interface{}) {
	l.writeLog(DEBUG, format, v...)
}

// Debugln logs a message string at the DEBUG level
func (l *logger) Debugln(v ...interface{}) {
	l.writeLog(DEBUG, fmt.Sprintln(v...))
}

// Info logs a message string at the INFO level
func (l *logger) Info(v ...interface{}) {
	l.writeLog(INFO, fmt.Sprint(v...))
}

// Infof logs a formatted message string at the INFO level
func (l *logger) Infof(format string, v ...interface{}) {
	l.writeLog(INFO, format, v...)
}

// Infoln logs a message string at the INFO level
func (l *logger) Infoln(v ...interface{}) {
	l.writeLog(INFO, fmt.Sprintln(v...))
}

// Warning logs a message string at the WARN level
func (l *logger) Warning(v ...interface{}) {
	l.writeLog(WARN, fmt.Sprint(v...))
}

// Warningf logs a formatted message string at the WARN level
func (l *logger) Warningf(format string, v ...interface{}) {
	l.writeLog(WARN, format, v...)
}

// Warningln logs a message string at the WARN level
func (l *logger) Warningln(v ...interface{}) {
	l.writeLog(WARN, fmt.Sprintln(v...))
}

// Error logs a message string at the ERROR level
func (l *logger) Error(v ...interface{}) {
	l.writeLog(ERROR, fmt.Sprint(v...))
}

// Errorf logs a formatted message string at the ERROR level
func (l *logger) Errorf(format string, v ...interface{}) {
	l.writeLog(ERROR, format, v...)
}

// Errorln logs a message string at the ERROR level
func (l *logger) Errorln(v ...interface{}) {
	l.writeLog(ERROR, fmt.Sprintln(v...))
}

// Fatal logs a message string at the FATAL level, and exits (via os.Exit)
// It overrides the golang standard library log.Fatal() method and prepends "FATAL" to the log message
// for consistency with other leveled logging methods in this package.
func (l *logger) Fatal(v ...interface{}) {
	l.writeLog(FATAL, fmt.Sprint(v...))

	if !l.test {
		os.Exit(1)
	}
}

// Fatalf logs a formatted message string at the FATAL level, and exits (via os.Exit)
// It overrides the golang standard library log.Fatal() method and prepends "FATAL" to the log message
// for consistency with other leveled logging methods in this package.
func (l *logger) Fatalf(format string, v ...interface{}) {
	l.writeLog(FATAL, format, v...)

	if !l.test {
		os.Exit(1)
	}
}

// Fatalln logs a message string at the FATAL level, and exits (via os.Exit)
// It overrides the golang standard library log.Fatal() method and prepends "FATAL" to the log message
// for consistency with other leveled logging methods in this package.
func (l *logger) Fatalln(v ...interface{}) {
	l.writeLog(FATAL, fmt.Sprint(v...))

	if !l.test {
		os.Exit(1)
	}
}

// Log is an implementation of the AwsLogger interface which will write an unleveled log message (meaning it will not
// be filtered by the log level checking logic, and is always written to the output writer)
func (l *logger) Log(v ...interface{}) {
	_ = l.Output(2, fmt.Sprint(v...))
}

func (l *logger) writeLog(lvl uint8, format string, v ...interface{}) {
	if l.level >= lvl {
		prefix := levels[lvl]
		_ = l.Output(3, prefix+" "+fmt.Sprintf(format, v...))
	}
}
