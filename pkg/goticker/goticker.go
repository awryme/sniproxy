package goticker

import (
	"time"
)

type Action func(elapsed time.Duration)

func Run(d time.Duration, onTick Action) (stop func()) {
	ticker := time.NewTicker(d)
	stopper := make(chan struct{})

	go func() {
		start := time.Now()
		for {
			select {
			case <-ticker.C:
				onTick(time.Since(start))
			case <-stopper:
				return
			}
		}
	}()

	return func() {
		ticker.Stop()
		close(stopper)
	}
}
