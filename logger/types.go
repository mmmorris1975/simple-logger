package logger

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
	// DEBUG log level outputs messages at the DEBUG level or higher
	DEBUG
)

// an array of numeric log levels (as array indices), mapped to the log level string representation.
var levels [DEBUG + 1]string

func init() {
	levels[NONE] = "NONE"
	levels[FATAL] = "FATAL"
	levels[ERROR] = "ERROR"
	levels[WARN] = "WARN"
	levels[INFO] = "INFO"
	levels[DEBUG] = "DEBUG"
}

// GoLogger is an interface describing the methods of the golang standard library log package to write log messages
type GoLogger interface {
	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})
	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})
}

// LeveledLogger is an interface which describes a leveled logging implementation, which may or may not leverage
// the standard library log package facilities
type LeveledLogger interface {
	Debug(...interface{})
	Debugf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Warning(...interface{})
	Warningf(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
}

// AwsLogger defines a logging implementation compatible with the AWS Go SDK Logger interface
type AwsLogger interface {
	Log(...interface{})
}
