package dogger

import (
	"bytes"
	"context"
	"fmt"
	logCtx "github.com/nikwo/dogger/context"
	"github.com/nikwo/dogger/format"
	"github.com/nikwo/dogger/level"
	"github.com/nikwo/dogger/utility"
	"io"
	"sync"
)

type Logger interface {
	WithContext(ctx context.Context) Logger
	WithFields(entry string, value interface{}) Logger
	Trace(input interface{})
	Debug(input interface{})
	Info(input interface{})
	Warn(input interface{})
	Error(input interface{})
}

type logger struct {
	lock      sync.Mutex
	lvl       level.Level
	ctx       logCtx.LogContext
	buffer    *bytes.Buffer
	formatter format.Format
	fields    map[string]interface{}
}

func newChildLogger() *logger {
	l := new(logger)
	lockIO()
	defer unlockIO()
	l.lvl = log.lvl
	l.buffer = bytes.NewBuffer([]byte{})
	l.formatter = log.formatter
	return l
}

func (l *logger) acceptedLevel(lvl level.Level) bool {
	return l.lvl <= lvl
}

func createChildLogger(ctx context.Context, lvl level.Level) *logger {
	l := newChildLogger()
	l.ctx = logCtx.NewLogContext(ctx, lvl, 1)

	return l
}

func (l *logger) WithContext(ctx context.Context) Logger {
	l.ctx = logCtx.NewLogContext(ctx, level.TRACE)
	return l
}

func (l *logger) WithFields(entry string, value interface{}) Logger {
	l.fields = make(map[string]interface{})
	l.fields[entry] = value
	return l
}

func (l *logger) Trace(input interface{}) {
	l.ctx.SetLevel(level.TRACE)
	if !l.acceptedLevel(level.TRACE) {
		return
	}

	l.formOutput(input)
	l.flush()
}

func (l *logger) Debug(input interface{}) {
	l.ctx.SetLevel(level.DEBUG)
	if !l.acceptedLevel(level.DEBUG) {
		return
	}

	l.formOutput(input)
	l.flush()
}

func (l *logger) Info(input interface{}) {
	l.ctx.SetLevel(level.INFO)
	if !l.acceptedLevel(level.INFO) {
		return
	}

	l.formOutput(input)
	l.flush()
}

func (l *logger) Warn(input interface{}) {
	l.ctx.SetLevel(level.WARN)
	if !l.acceptedLevel(level.WARN) {
		return
	}

	l.formOutput(input)
	l.flush()
}

func (l *logger) Error(input interface{}) {
	l.ctx.SetLevel(level.ERROR)

	if !l.acceptedLevel(level.ERROR) {
		return
	}

	l.formOutput(input)
	l.flush()
}

func (l *logger) formOutput(input interface{}) {
	inlineBuffer := make([]byte, 0, 100)
	inlineBuffer = append(inlineBuffer, utility.Bytes(l.formatter.FormatString(l.ctx))...)
	for entry, field := range l.fields {
		inlineBuffer = append(inlineBuffer, utility.Bytes(fmt.Sprintf(" entry=\"%s\" value=\"%+v\"", entry, field))...)
	}
	inlineBuffer = append(inlineBuffer, utility.Bytes(fmt.Sprintf(" message=\"%+v\"\n", input))...)
	_, _ = l.buffer.Write(inlineBuffer)
}

func (l *logger) flush() {
	lockIO()
	defer unlockIO()
	_, _ = writer.Write(l.buffer.Bytes())
}

var (
	log    *logger
	writer io.Writer
)

func lockIO() {
	log.lock.Lock()
}

func unlockIO() {
	log.lock.Unlock()
}
