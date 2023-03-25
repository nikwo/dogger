package dogger

import (
	"context"
	"github.com/nikwo/dogger/format"
	"io"

	"github.com/nikwo/dogger/level"
)

func WithContext(ctx context.Context) Logger {
	l := createChildLogger(ctx, level.TRACE)
	return l
}

func WithFields(entry string, value interface{}) Logger {
	l := createChildLogger(context.Background(), level.TRACE)
	l.fields = make(map[string]interface{})
	l.fields[entry] = value
	return l
}

func Trace(input interface{}) {
	l := createChildLogger(context.Background(), level.TRACE)
	if !l.acceptedLevel(level.TRACE) {
		return
	}

	l.formOutput(input)
	l.flush()
}

func Debug(input interface{}) {
	l := createChildLogger(context.Background(), level.DEBUG)
	if !l.acceptedLevel(level.DEBUG) {
		return
	}

	l.formOutput(input)
	l.flush()
}

func Info(input interface{}) {
	l := createChildLogger(context.Background(), level.INFO)
	if !l.acceptedLevel(level.INFO) {
		return
	}

	l.formOutput(input)
	l.flush()
}

func Warn(input interface{}) {
	l := createChildLogger(context.Background(), level.WARN)
	if !l.acceptedLevel(level.WARN) {
		return
	}

	l.formOutput(input)
	l.flush()
}

func Error(input interface{}) {
	l := createChildLogger(context.Background(), level.ERROR)
	if !l.acceptedLevel(level.ERROR) {
		return
	}

	l.formOutput(input)
	l.flush()
}

func SetLevel(level level.Level) {
	lockIO()
	defer unlockIO()
	log.lvl = level
}

func SetWriter(customWriter io.Writer) {
	lockIO()
	defer unlockIO()
	writer = customWriter
}

func SetFormatter(formatter format.Format) {
	lockIO()
	defer unlockIO()
	log.formatter = formatter
}

func ExportDefaultLogger() Logger {
	return log
}
