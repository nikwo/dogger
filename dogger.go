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
	"os"
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

func WithContext(ctx context.Context) Logger {
	l := createChildLogger(ctx, level.TRACE)
	return l
}

func (l *logger) WithContext(ctx context.Context) Logger {
	l.ctx = logCtx.NewLogContext(ctx, level.TRACE)
	return l
}

func WithFields(entry string, value interface{}) Logger {
	l := createChildLogger(context.Background(), level.TRACE)
	l.fields = make(map[string]interface{})
	l.fields[entry] = value
	return l
}

func (l *logger) WithFields(entry string, value interface{}) Logger {
	l.fields = make(map[string]interface{})
	l.fields[entry] = value
	return l
}

func Trace(input interface{}) {
	l := createChildLogger(context.Background(), level.TRACE)
	if !l.acceptedLevel(level.TRACE) {
		return
	}

	l.FormOutput(input)
	l.Flush()
}

func (l *logger) Trace(input interface{}) {
	l.ctx.SetLevel(level.TRACE)
	if !l.acceptedLevel(level.TRACE) {
		return
	}

	l.FormOutput(input)
	l.Flush()
}

func Debug(input interface{}) {
	l := createChildLogger(context.Background(), level.DEBUG)
	if !l.acceptedLevel(level.DEBUG) {
		return
	}

	l.FormOutput(input)
	l.Flush()
}

func (l *logger) Debug(input interface{}) {
	l.ctx.SetLevel(level.DEBUG)
	if !l.acceptedLevel(level.DEBUG) {
		return
	}

	l.FormOutput(input)
	l.Flush()
}

func Info(input interface{}) {
	l := createChildLogger(context.Background(), level.INFO)
	if !l.acceptedLevel(level.INFO) {
		return
	}

	l.FormOutput(input)
	l.Flush()
}

func (l *logger) Info(input interface{}) {
	l.ctx.SetLevel(level.INFO)
	if !l.acceptedLevel(level.INFO) {
		return
	}

	l.FormOutput(input)
	l.Flush()
}

func Warn(input interface{}) {
	l := createChildLogger(context.Background(), level.WARN)
	if !l.acceptedLevel(level.WARN) {
		return
	}

	l.FormOutput(input)
	l.Flush()
}

func (l *logger) Warn(input interface{}) {
	l.ctx.SetLevel(level.WARN)
	if !l.acceptedLevel(level.WARN) {
		return
	}

	l.FormOutput(input)
	l.Flush()
}

func Error(input interface{}) {
	l := createChildLogger(context.Background(), level.ERROR)
	if !l.acceptedLevel(level.ERROR) {
		return
	}

	l.FormOutput(input)
	l.Flush()
}

func (l *logger) Error(input interface{}) {
	l.ctx.SetLevel(level.ERROR)

	if !l.acceptedLevel(level.ERROR) {
		return
	}

	l.FormOutput(input)
	l.Flush()
}

func (l *logger) FormOutput(input interface{}) {
	inlineBuffer := make([]byte, 0, 100)
	inlineBuffer = append(inlineBuffer, utility.Bytes(l.formatter.FormatString(l.ctx))...)
	for entry, field := range l.fields {
		inlineBuffer = append(inlineBuffer, utility.Bytes(fmt.Sprintf(" entry=\"%s\" value=\"%+v\"", entry, field))...)
	}
	inlineBuffer = append(inlineBuffer, utility.Bytes(fmt.Sprintf(" message=\"%+v\"\n", input))...)
	_, _ = l.buffer.Write(inlineBuffer)
}

func (l *logger) Flush() {
	lockIO()
	defer unlockIO()
	_, _ = writer.Write(l.buffer.Bytes())
}

var (
	log    *logger
	writer io.Writer
)

func init() {
	log = &logger{
		lvl:       level.TRACE,
		formatter: format.DefaultFormatter(),
	}
	writer = os.Stdout
}

func lockIO() {
	log.lock.Lock()
}

func unlockIO() {
	log.lock.Unlock()
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
