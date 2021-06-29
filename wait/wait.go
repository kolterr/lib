package wait

import (
	"sync"
	"time"
)

type Wait struct {
	sync.WaitGroup
}

func (w *Wait) Wrap(fn func()) {
	w.Add(1)
	go func() {
		fn()
		w.Done()
	}()
}

// WaitWithTimeout blocks until the WaitGroup counter is zero or timeout
// returns true if timeout
func (w *Wait) WaitWithTimeout(t time.Duration) bool {
	ch := make(chan struct{})
	go func() {
		defer close(ch)
		w.Wait()
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		return false
	case <-time.After(t):
		return true
	}
}
