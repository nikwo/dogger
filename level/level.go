package level

import (
	"github.com/nikwo/dogger/errors"
)

const (
	TRACE = iota
	DEBUG
	INFO
	WARN
	ERROR
)

type Level int

func (l Level) String() string {
	switch l {
	case TRACE:
		return "trace"
	case DEBUG:
		return "debug"
	case INFO:
		return "info"
	case WARN:
		return "warn"
	case ERROR:
		return "error"
	default:
		return "info"
	}
}

func LogLevelFromString(level string) (Level, error) {
	switch level {
	case "trace":
		return TRACE, nil
	case "debug":
		return DEBUG, nil
	case "info":
		return INFO, nil
	case "warn":
		return WARN, nil
	case "error":
		return ERROR, nil
	default:
		return INFO, errors.ErrInvalidLevelString
	}
}
