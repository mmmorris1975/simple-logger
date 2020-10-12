package logger

import (
	"log"
	"os"
	"testing"
)

// Return a new logger for use with Example tests, sending output to stdout, also
// with no flags which could potentially change the output across test runs
func newLogger(level uint8) *logger {
	return &logger{
		Logger: log.New(os.Stdout, "", 0),
		level:  level,
	}
}

func nullLogger() *logger {
	f, _ := os.Open(os.DevNull)
	return &logger{
		Logger: log.New(f, "", 0),
		level:  NONE,
	}
}

func TestParseLevel(t *testing.T) {
	t.Run("good", func(t *testing.T) {
		for i, v := range levels {
			lvl, err := ParseLevel(v)
			if err != nil {
				t.Error("err")
				continue
			}

			if i != int(lvl) {
				t.Error("data mismatch")
			}
		}
	})

	t.Run("invalid", func(t *testing.T) {
		if _, err := ParseLevel("invalid"); err == nil {
			t.Error("did not receive expected error")
		}
	})
}

func TestLogger_SetLevel(t *testing.T) {
	l := nullLogger()
	l.SetLevel(DEBUG)

	if l.level != DEBUG {
		t.Error("data mismatch")
	}
}

func TestLogger_WithLevel(t *testing.T) {
	l := nullLogger().WithLevel(DEBUG)

	if l.level != DEBUG {
		t.Error("data mismatch")
	}
}

// No testing of Fatal*() logging, since it calls os.Exit in stdlib log.Fatal*()
func TestLogger_Panic(t *testing.T) {
	defer func() {
		if x := recover(); x == nil {
			t.Errorf("Did not receive expected panic")
		}
	}()

	nullLogger().Panic("oh no!")
}

func Example_logger_Debug_on() {
	newLogger(DEBUG).Debug("debug log message")
	// Output:
	// DEBUG debug log message
}

func Example_logger_Debug_off() {
	newLogger(INFO).Debug("debug log message")
	// Output:
	//
}

func Example_logger_Error_on() {
	newLogger(ERROR).Error("error log message")
	// Output:
	// ERROR error log message
}

func Example_logger_Error_off() {
	newLogger(NONE).Error("error log message")
	// Output:
	//
}

func Example_logger_Info_on() {
	newLogger(INFO).Info("info log message")
	// Output:
	// INFO info log message
}

func Example_logger_Info_off() {
	newLogger(FATAL).Info("info log message")
	// Output:
	//
}

func Example_logger_Warning_on() {
	newLogger(WARN).Warning("warning log message")
	// Output:
	// WARN warning log message
}

func Example_logger_Warning_off() {
	newLogger(ERROR).Warning("warning log message")
	// Output:
	//
}

func ExampleLogger_Print() {
	newLogger(NONE).Print("Print() log message")
	// Output:
	// Print() log message
}

func Example_logger_Log() {
	newLogger(NONE).Log("Log() log message")
	// Output:
	// Log() log message
}
