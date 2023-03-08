package format

import (
	"fmt"
	"time"

	"github.com/fatih/color"
	"github.com/nikwo/dogger/context"
)

const (
	DefaultTimeFormat = time.RFC3339
)

var colorizers = map[string]func(...interface{}) string{
	"trace":   color.New(color.BgBlack).Add(color.FgHiWhite).SprintFunc(),
	"debug":   color.New(color.BgBlack).Add(color.FgHiBlue).SprintFunc(),
	"info":    color.New(color.BgBlack).Add(color.FgHiGreen).SprintFunc(),
	"warn":    color.New(color.BgBlack).Add(color.FgHiYellow).SprintFunc(),
	"error":   color.New(color.BgBlack).Add(color.FgHiRed).SprintFunc(),
	"default": color.New(color.BgBlack).Add(color.FgWhite).SprintFunc(),
}

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
		logContext.GetCaller(),
	)
}

func (df *defaultFormat) MatchVerboseColor(lvl string) string {
	colorizer, exists := colorizers[lvl]
	if !exists {
		return colorizers["default"](lvl)
	}

	return colorizer(lvl)
}

func DefaultFormatter() Format {
	return &defaultFormat{}
}
