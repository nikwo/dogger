package main

import (
	"context"
	"github.com/enescakir/emoji"

	"github.com/nikwo/dogger"
)

func main() {
	ctx := context.Background()
	dogger.WithContext(ctx).Trace("hello")
	dogger.Debug("world")
	dogger.WithContext(ctx).WithFields("chiba", "inu").Info(emoji.Sprintf("i :heart:  pugs!"))
	dogger.Info("dogs aren't cute")
	dogger.WithContext(ctx).Error("error: dogs are cute, you've lied!")
}
