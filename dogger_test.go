package dogger

import (
	"github.com/nikwo/dogger/level"
	"testing"
)

func TestLevelTrace(t *testing.T) {
	result, err := level.LogLevelFromString("trace")
	if err != nil {
		t.Fail()
	}
	if result != level.TRACE {
		t.Fail()
	}
}

func TestLevelDebug(t *testing.T) {
	result, err := level.LogLevelFromString("debug")
	if err != nil {
		t.Fail()
	}
	if result != level.DEBUG {
		t.Fail()
	}
}

func TestLevelInfo(t *testing.T) {
	result, err := level.LogLevelFromString("info")
	if err != nil {
		t.Fail()
	}
	if result != level.INFO {
		t.Fail()
	}
}

func TestLevelWarn(t *testing.T) {
	result, err := level.LogLevelFromString("warn")
	if err != nil {
		t.Fail()
	}
	if result != level.WARN {
		t.Fail()
	}
}

func TestLevelError(t *testing.T) {
	result, err := level.LogLevelFromString("error")
	if err != nil {
		t.Fail()
	}
	if result != level.ERROR {
		t.Fail()
	}
}

func TestDogger(t *testing.T) {
	SetLevel(level.TRACE)
	Trace(struct {
		Hello string
	}{Hello: "World"})
}
