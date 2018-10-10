package simple_logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

const (
	NONE = iota
	FATAL
	ERROR
	WARN
	INFO
	DEBUG
)

var levels = []string{"NONE", "FATAL", "ERROR", "WARN", "INFO", "DEBUG"}

type Logger struct {
	level uint
	*log.Logger
}

var StdLogger = NewLogger(os.Stderr, "", log.LstdFlags)

func NewLogger(out io.Writer, prefix string, flag int) *Logger {
	l := &Logger{level: INFO, Logger: log.New(out, prefix, flag)}
	return l
}

func (l *Logger) SetLevel(level uint) {
	l.level = level
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.writeLogf(FATAL, format, v...)
	os.Exit(1)
}

func (l *Logger) Fatal(v ...interface{}) {
	l.writeLogln(FATAL, v...)
	os.Exit(1)
}

func (l *Logger) Fatalln(v ...interface{}) {
	l.writeLogln(FATAL, v...)
	os.Exit(1)
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.writeLogf(ERROR, format, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.writeLogln(ERROR, v...)
}

func (l *Logger) Errorln(v ...interface{}) {
	l.writeLogln(ERROR, v...)
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.writeLogf(WARN, format, v...)
}

func (l *Logger) Warn(v ...interface{}) {
	l.writeLogln(WARN, v...)
}

func (l *Logger) Warnln(v ...interface{}) {
	l.writeLogln(WARN, v...)
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.writeLogf(INFO, format, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.writeLogln(INFO, v...)
}

func (l *Logger) Infoln(v ...interface{}) {
	l.writeLogln(INFO, v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.writeLogf(DEBUG, format, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	l.writeLogln(DEBUG, v...)
}

func (l *Logger) Debugln(v ...interface{}) {
	l.writeLogln(DEBUG, v...)
}

func (l *Logger) writeLogf(level uint, format string, v ...interface{}) {
	if l.level >= level {
		q := []interface{}{levels[level]}
		q = append(q, v...)
		l.Output(3, fmt.Sprintf("%s "+format, q...))
	}
}

func (l *Logger) writeLogln(level uint, v ...interface{}) {
	if l.level >= level {
		q := []interface{}{levels[level]}
		q = append(q, v...)
		l.Output(3, fmt.Sprintln(q...))
	}
}
