package dogger

import (
	"testing"

	"github.com/nikwo/dogger/level"
)

func TestLevelTrace(t *testing.T) {
	result := level.LogLevelFromString("trace")
	if result != level.TRACE {
		t.Fail()
	}
}

func TestLevelDebug(t *testing.T) {
	result := level.LogLevelFromString("debug")
	if result != level.DEBUG {
		t.Fail()
	}
}

func TestLevelInfo(t *testing.T) {
	result := level.LogLevelFromString("info")
	if result != level.INFO {
		t.Fail()
	}
}

func TestLevelWarn(t *testing.T) {
	result := level.LogLevelFromString("warn")
	if result != level.WARN {
		t.Fail()
	}
}

func TestLevelError(t *testing.T) {
	result := level.LogLevelFromString("error")
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
