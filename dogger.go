package dogger

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	logCtx "github.com/nikwo/dogger/context"
	"github.com/nikwo/dogger/format"
	"github.com/nikwo/dogger/level"
)

type Logger interface {
	WithContext(ctx context.Context) Logger
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
}

func newChildLogger() *logger {
	l := new(logger)
	log.lock.Lock()
	defer log.lock.Unlock()
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
	l.ctx = logCtx.NewLogContext(ctx, lvl)

	return l
}

func WithContext(ctx context.Context) Logger {
	l := newChildLogger()
	l.ctx = logCtx.NewLogContext(ctx, level.TRACE)
	return l
}

func (l *logger) WithContext(ctx context.Context) Logger {
	l.ctx = logCtx.NewLogContext(ctx, level.TRACE)
	return l
}

func Trace(input interface{}) {
	l := createChildLogger(context.Background(), level.TRACE)
	if !l.acceptedLevel(level.TRACE) {
		return
	}

	inlineBuffer := ([]byte)(l.formatter.FormatString(l.ctx) + fmt.Sprintf(" %+v\n", input))
	messages <- inlineBuffer
}

func (l *logger) Trace(input interface{}) {
	l.ctx.SetLevel(level.TRACE)
	if !l.acceptedLevel(level.TRACE) {
		return
	}

	_, _ = l.buffer.Write(([]byte)(l.formatter.FormatString(l.ctx) + fmt.Sprintf(" %+v\n", input)))
	messages <- l.buffer.Bytes()
}

func Debug(input interface{}) {
	l := createChildLogger(context.Background(), level.DEBUG)
	if !l.acceptedLevel(level.DEBUG) {
		return
	}

	inlineBuffer := ([]byte)(l.formatter.FormatString(l.ctx) + fmt.Sprintf(" %+v\n", input))
	messages <- inlineBuffer
}

func (l *logger) Debug(input interface{}) {
	l.ctx.SetLevel(level.DEBUG)
	if !l.acceptedLevel(level.DEBUG) {
		return
	}

	_, _ = l.buffer.Write(([]byte)(l.formatter.FormatString(l.ctx) + fmt.Sprintf(" %+v\n", input)))
	messages <- l.buffer.Bytes()
}

func Info(input interface{}) {
	l := createChildLogger(context.Background(), level.INFO)
	if !l.acceptedLevel(level.INFO) {
		return
	}

	inlineBuffer := ([]byte)(l.formatter.FormatString(l.ctx) + fmt.Sprintf(" %+v\n", input))
	messages <- inlineBuffer
}

func (l *logger) Info(input interface{}) {
	l.ctx.SetLevel(level.INFO)
	if !l.acceptedLevel(level.INFO) {
		return
	}

	_, _ = l.buffer.Write(([]byte)(l.formatter.FormatString(l.ctx) + fmt.Sprintf(" %+v\n", input)))
	messages <- l.buffer.Bytes()
}

func Warn(input interface{}) {
	l := createChildLogger(context.Background(), level.WARN)
	if !l.acceptedLevel(level.WARN) {
		return
	}

	inlineBuffer := ([]byte)(l.formatter.FormatString(l.ctx) + fmt.Sprintf(" %+v\n", input))
	messages <- inlineBuffer
}

func (l *logger) Warn(input interface{}) {
	l.ctx.SetLevel(level.WARN)
	if !l.acceptedLevel(level.WARN) {
		return
	}

	_, _ = l.buffer.Write(([]byte)(l.formatter.FormatString(l.ctx) + fmt.Sprintf(" %+v\n", input)))
	messages <- l.buffer.Bytes()
}

func Error(input interface{}) {
	l := createChildLogger(context.Background(), level.ERROR)
	if !l.acceptedLevel(level.ERROR) {
		return
	}

	inlineBuffer := ([]byte)(l.formatter.FormatString(l.ctx) + fmt.Sprintf(" %+v\n", input))
	messages <- inlineBuffer
}

func (l *logger) Error(input interface{}) {
	l.ctx.SetLevel(level.ERROR)

	if !l.acceptedLevel(level.ERROR) {
		return
	}

	_, _ = l.buffer.Write(([]byte)(l.formatter.FormatString(l.ctx) + fmt.Sprintf(" %+v\n", input)))
	messages <- l.buffer.Bytes()
}

var (
	log = &logger{
		lvl:       level.TRACE,
		formatter: format.DefaultFormatter(),
	}
	writer   = os.Stdout
	messages = make(chan []byte, 100)
)

func init() {
	go background(context.Background())
}

func background(ctx context.Context) {
	for {
		select {
		case buffer := <-messages:
			_, _ = writer.Write(buffer)
		case <-ctx.Done():
			for len(messages) > 0 {
				<-time.Tick(time.Millisecond * 100)
			}
			return
		}
	}
}

func SetLevel(level level.Level) {
	log.lock.Lock()
	defer log.lock.Unlock()
	log.lvl = level
}
