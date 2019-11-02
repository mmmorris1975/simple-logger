package logger

import (
	"bytes"
	"log"
	"os"
	"testing"
)

// No testing of Fatal*() logging, since it calls os.Exit
func TestNewLogger(t *testing.T) {
	t.Run("std logger", func(t *testing.T) {
		l := StdLogger
		if l.Level != INFO {
			t.Error("Level mismatch")
		}

		if l.Flags() != log.LstdFlags {
			t.Error("flag mismatch")
		}
	})

	t.Run("debug logger", func(t *testing.T) {
		l := NewLogger(os.Stderr, "", log.LstdFlags)
		l.SetLevel(DEBUG)
		if l.Level != DEBUG {
			t.Error("Level mismatch")
		}
	})

	t.Run("flags", func(t *testing.T) {
		l := NewLogger(os.Stderr, "", log.Lshortfile|log.Lmicroseconds)
		if l.Flags() != log.Lshortfile|log.Lmicroseconds {
			t.Error("flag mismatch")
		}
	})
}

func TestLogger_Panic(t *testing.T) {
	t.Run("panic", func(t *testing.T) {
		defer func() {
			if x := recover(); x == nil {
				t.Error("did not receive expected panic")
			}
		}()
		l := exampleLogger(NONE)
		l.Panic("halp!")
	})

	t.Run("panicf", func(t *testing.T) {
		defer func() {
			if x := recover(); x == nil {
				t.Error("did not receive expected panic")
			}
		}()
		l := exampleLogger(NONE)
		l.Panicf("halp!")
	})

	t.Run("panicln", func(t *testing.T) {
		defer func() {
			if x := recover(); x == nil {
				t.Error("did not receive expected panic")
			}
		}()
		l := exampleLogger(NONE)
		l.Panicln("halp!")
	})
}

func TestLogger_Debug(t *testing.T) {
	b := new(bytes.Buffer)
	l := StdLogger
	l.SetOutput(b)
	l.SetLevel(DEBUG)

	t.Run("debug", func(t *testing.T) {
		b.Reset()
		l.Debug("test")
		if b.Len() < 1 {
			t.Error("expected output to be logged")
		}
	})

	t.Run("info", func(t *testing.T) {
		b.Reset()
		l.Info("test")
		if b.Len() < 1 {
			t.Error("expected output to be logged")
		}
	})

	t.Run("warn", func(t *testing.T) {
		b.Reset()
		l.Warn("test")
		if b.Len() < 1 {
			t.Error("expected output to be logged")
		}
	})

	t.Run("error", func(t *testing.T) {
		b.Reset()
		l.Error("test")
		if b.Len() < 1 {
			t.Error("expected output to be logged")
		}
	})
}

func TestLogger_Info(t *testing.T) {
	b := new(bytes.Buffer)
	l := StdLogger
	l.SetOutput(b)
	l.SetLevel(INFO)

	t.Run("debug", func(t *testing.T) {
		b.Reset()
		l.Debug("test")
		if b.Len() > 1 {
			t.Error("expected output to not be logged")
		}
	})

	t.Run("info", func(t *testing.T) {
		b.Reset()
		l.Info("test")
		if b.Len() < 1 {
			t.Error("expected output to be logged")
		}
	})

	t.Run("warn", func(t *testing.T) {
		b.Reset()
		l.Warn("test")
		if b.Len() < 1 {
			t.Error("expected output to be logged")
		}
	})

	t.Run("error", func(t *testing.T) {
		b.Reset()
		l.Error("test")
		if b.Len() < 1 {
			t.Error("expected output to be logged")
		}
	})
}

func TestLogger_Warn(t *testing.T) {
	b := new(bytes.Buffer)
	l := StdLogger
	l.SetOutput(b)
	l.SetLevel(WARN)

	t.Run("debug", func(t *testing.T) {
		b.Reset()
		l.Debug("test")
		if b.Len() > 1 {
			t.Error("expected output to not be logged")
		}
	})

	t.Run("info", func(t *testing.T) {
		b.Reset()
		l.Info("test")
		if b.Len() > 1 {
			t.Error("expected output to not be logged")
		}
	})

	t.Run("warn", func(t *testing.T) {
		b.Reset()
		l.Warn("test")
		if b.Len() < 1 {
			t.Error("expected output to be logged")
		}
	})

	t.Run("error", func(t *testing.T) {
		b.Reset()
		l.Error("test")
		if b.Len() < 1 {
			t.Error("expected output to be logged")
		}
	})
}

func TestLogger_Error(t *testing.T) {
	b := new(bytes.Buffer)
	l := StdLogger
	l.SetOutput(b)
	l.SetLevel(ERROR)

	t.Run("debug", func(t *testing.T) {
		b.Reset()
		l.Debug("test")
		if b.Len() > 1 {
			t.Error("expected output to not be logged")
		}
	})

	t.Run("info", func(t *testing.T) {
		b.Reset()
		l.Info("test")
		if b.Len() > 1 {
			t.Error("expected output to not be logged")
		}
	})

	t.Run("warn", func(t *testing.T) {
		b.Reset()
		l.Warn("test")
		if b.Len() > 1 {
			t.Error("expected output to not be logged")
		}
	})

	t.Run("error", func(t *testing.T) {
		b.Reset()
		l.Error("test")
		if b.Len() < 1 {
			t.Error("expected output to be logged")
		}
	})
}

func TestLogger_None(t *testing.T) {
	b := new(bytes.Buffer)
	l := StdLogger
	l.SetOutput(b)
	l.SetLevel(NONE)

	t.Run("debug", func(t *testing.T) {
		b.Reset()
		l.Debug("test")
		if b.Len() > 1 {
			t.Error("expected output to not be logged")
		}
	})

	t.Run("info", func(t *testing.T) {
		b.Reset()
		l.Info("test")
		if b.Len() > 1 {
			t.Error("expected output to not be logged")
		}
	})

	t.Run("warn", func(t *testing.T) {
		b.Reset()
		l.Warn("test")
		if b.Len() > 1 {
			t.Error("expected output to not be logged")
		}
	})

	t.Run("error", func(t *testing.T) {
		b.Reset()
		l.Error("test")
		if b.Len() > 1 {
			t.Error("expected output to not be logged")
		}
	})
}

func TestParseLevel(t *testing.T) {
	t.Run("none", func(t *testing.T) {
		l, err := ParseLevel("None")
		if err != nil {
			t.Error(err)
			return
		}
		if l != NONE {
			t.Errorf("Level mismatch: EXPECTED %d, GOT %d", NONE, l)
		}
	})

	t.Run("fatal", func(t *testing.T) {
		l, err := ParseLevel("fatal")
		if err != nil {
			t.Error(err)
			return
		}
		if l != FATAL {
			t.Errorf("Level mismatch: EXPECTED %d, GOT %d", FATAL, l)
		}
	})

	t.Run("error", func(t *testing.T) {
		l, err := ParseLevel("ERROR")
		if err != nil {
			t.Error(err)
			return
		}
		if l != ERROR {
			t.Errorf("Level mismatch: EXPECTED %d, GOT %d", ERROR, l)
		}
	})

	t.Run("warn", func(t *testing.T) {
		l, err := ParseLevel("wArN")
		if err != nil {
			t.Error(err)
			return
		}
		if l != WARN {
			t.Errorf("Level mismatch: EXPECTED %d, GOT %d", WARN, l)
		}
	})

	t.Run("info", func(t *testing.T) {
		l, err := ParseLevel("InFo")
		if err != nil {
			t.Error(err)
			return
		}
		if l != INFO {
			t.Errorf("Level mismatch: EXPECTED %d, GOT %d", INFO, l)
		}
	})

	t.Run("debug", func(t *testing.T) {
		l, err := ParseLevel("DebuG")
		if err != nil {
			t.Error(err)
			return
		}
		if l != DEBUG {
			t.Errorf("Level mismatch: EXPECTED %d, GOT %d", DEBUG, l)
		}
	})

	t.Run("invalid", func(t *testing.T) {
		if _, err := ParseLevel("bogus"); err == nil {
			t.Error("did not receive expected error")
			return
		}
	})
}

func ExampleLogger_Debug() {
	l := exampleLogger(DEBUG)
	l.Debug("test")
	// Output:
	// DEBUG test
}

func ExampleLogger_Debugf() {
	l := exampleLogger(DEBUG)
	l.Debugf("%s", "test")
	// Output:
	// DEBUG test
}

func ExampleLogger_Debugln() {
	l := exampleLogger(DEBUG)
	l.Debugln("test")
	// Output:
	// DEBUG test
}

func ExampleLogger_Info() {
	l := exampleLogger(INFO)
	l.Info("test")
	// Output:
	// INFO test
}

func ExampleLogger_Infof() {
	l := exampleLogger(INFO)
	l.Infof("%s", "test")
	// Output:
	// INFO test
}

func ExampleLogger_Infoln() {
	l := exampleLogger(INFO)
	l.Infoln("test")
	// Output:
	// INFO test
}

func ExampleLogger_Warn() {
	l := exampleLogger(WARN)
	l.Warn("test")
	// Output:
	// WARN test
}

func ExampleLogger_Warnf() {
	l := exampleLogger(WARN)
	l.Warnf("%s", "test")
	// Output:
	// WARN test
}

func ExampleLogger_Warnln() {
	l := exampleLogger(WARN)
	l.Warnln("test")
	// Output:
	// WARN test
}

func ExampleLogger_Error() {
	l := exampleLogger(ERROR)
	l.Error("test")
	// Output:
	// ERROR test
}

func ExampleLogger_Errorf() {
	l := exampleLogger(ERROR)
	l.Errorf("%s", "test")
	// Output:
	// ERROR test
}

func ExampleLogger_Errorln() {
	l := exampleLogger(ERROR)
	l.Errorln("test")
	// Output:
	// ERROR test
}

func ExampleLogger_Panic() {
	// Panic*() logs should output something even if Level == NONE
	defer func() {
		if x := recover(); x != nil {
			// drop panic on the floor
		}
	}()
	l := exampleLogger(NONE)
	l.Panic("test")
	// Output:
	// PANIC test
}

func ExampleLogger_Panicf() {
	defer func() {
		if x := recover(); x != nil {
			// drop panic on the floor
		}
	}()
	l := exampleLogger(NONE)
	l.Panicf("%s", "test")
	// Output:
	// PANIC test
}

func ExampleLogger_Panicln() {
	defer func() {
		if x := recover(); x != nil {
			// drop panic on the floor
		}
	}()
	l := exampleLogger(NONE)
	l.Panicln("test")
	// Output:
	// PANIC test
}

func ExampleLogger_Print() {
	l := exampleLogger(INFO)
	l.Print("test")
	// Output:
	// test
}

func ExampleLogger_Printf() {
	l := exampleLogger(WARN)
	l.Printf("%s", "test")
	// Output:
	// test
}

func ExampleLogger_Println() {
	l := exampleLogger(ERROR)
	l.Println("test")
	// Output:
	// test
}

func ExampleLogger_Logf() {
	l := exampleLogger(WARN)
	l.Logf("%s %d", "test", 1)
	// Output:
	// test 1
}

func ExampleLogger_Log() {
	l := exampleLogger(ERROR)
	l.Log("test")
	// Output:
	// test
}

func ExampleLogger_Prefix() {
	l := exampleLogger(INFO)
	l.SetPrefix("123")
	l.Info("test")
	// Output:
	// 123INFO test
}

func exampleLogger(level uint) *Logger {
	// A logger for use with Example tests, sending output to stdout, also
	// with no flags which could potentially change the output across test runs
	l := NewLogger(os.Stdout, "", 0).WithLevel(level)
	return l
}