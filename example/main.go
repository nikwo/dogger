package main

import (
	"context"
	"sync"

	"github.com/nikwo/dogger"
)

func main() {
	dogger.WithContext(context.Background()).Trace("Hello")
	dogger.Debug("World")
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(counter int) {
			dogger.Info(counter)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
