package main

import (
	"context"
	"github.com/nikwo/dogger"
	"sync"
)

func main() {
	dogger.WithContext(context.Background()).Trace("Hello")
	dogger.Debug("World")
	wg := sync.WaitGroup{}
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func(counter int, wg *sync.WaitGroup) {
			dogger.Info(counter)
			wg.Done()
		}(i, &wg)
	}
	wg.Wait()
}
