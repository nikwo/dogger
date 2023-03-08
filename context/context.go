package context

import (
	"context"
	"runtime"
	"time"

	"github.com/nikwo/dogger/level"
	"github.com/nikwo/dogger/utility"
)

type LogContext interface {
	GetTime() time.Time
	GetCaller() string
	GetLevel() level.Level
	SetTime(t time.Time)
	SetCaller(caller string)
	SetLevel(lvl level.Level)
	Deadline() (deadline time.Time, ok bool)
	Done() <-chan struct{}
	Err() error
	Value(key interface{}) interface{}
}

func NewLogContext(ctx context.Context, lvl level.Level) LogContext {
	var caller string
	lc := &logContext{Context: ctx, Time: time.Now()}
	pc, _, _, ok := runtime.Caller(2)

	if !ok {
		return lc
	}

	details := runtime.FuncForPC(pc)
	file, line := details.FileLine(pc)
	caller = utility.FormatCaller(file, details.Name(), line)
	lc.Caller = caller
	lc.Level = lvl

	return lc
}

type logContext struct {
	Time   time.Time
	Caller string
	Level  level.Level
	context.Context
}

func (c *logContext) GetTime() time.Time {
	return c.Time
}

func (c *logContext) GetCaller() string {
	return c.Caller
}

func (c *logContext) GetLevel() level.Level {
	return c.Level
}

func (c *logContext) SetTime(t time.Time) {
	c.Time = t
}

func (c *logContext) SetCaller(caller string) {
	c.Caller = caller
}

func (c *logContext) SetLevel(lvl level.Level) {
	c.Level = lvl
}
