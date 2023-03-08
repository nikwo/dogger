package format

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/nikwo/dogger/context"
	"time"
)

const (
	DefaultTimeFormat = time.RFC3339
)

type Format interface {
	FormatString(ctx context.LogContext) string
	MatchVerboseColor(lvl string) string
}

type defaultFormat struct {
}

func (df *defaultFormat) FormatString(logContext context.LogContext) string {
	return fmt.Sprintf(
		"[%s] %s %s",
		df.MatchVerboseColor(logContext.GetLevel().String()),
		logContext.GetTime().Format(DefaultTimeFormat),
		logContext.GetCaller())
}

func (df *defaultFormat) MatchVerboseColor(lvl string) string {
	var colorize func(a ...interface{}) string
	switch lvl {
	case "trace":
		colorize = color.New(color.BgBlack).Add(color.FgHiWhite).SprintFunc()
	case "debug":
		colorize = color.New(color.BgBlack).Add(color.FgHiBlue).SprintFunc()
	case "info":
		colorize = color.New(color.BgBlack).Add(color.FgHiGreen).SprintFunc()
	case "warn":
		colorize = color.New(color.BgBlack).Add(color.FgHiYellow).SprintFunc()
	case "error":
		colorize = color.New(color.BgBlack).Add(color.FgHiRed).SprintFunc()
	default:
		colorize = color.New(color.BgBlack).Add(color.FgWhite).SprintFunc()
	}
	return colorize(lvl)
}

func DefaultFormatter() Format {
	return &defaultFormat{}
}
