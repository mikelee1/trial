package ratelimit

import (
	"fmt"
	"sync"
)

var limiter = &limitListener{}

var once = sync.Once{}

func NewLimit(n int) *limitListener {
	once.Do(func() {
		limiter = &limitListener{
			sem:  make(chan struct{}, n),
			done: make(chan struct{}),
		}

	})
	return limiter
}

func (l *limitListener) Wait() {
	done := false
	fmt.Println(cap(l.sem) - len(l.sem))
	select {
	case <-l.done:
		done = true
	case l.sem <- struct{}{}:
	}
	if done {
		l.closeOnce.Do(
			func() {
				close(l.done)
			},
		)
	}
}

func (l *limitListener) Release() {
	l.release()
}

type limitListener struct {
	lock      sync.RWMutex
	sem       chan struct{}
	closeOnce sync.Once     // ensures the done chan is only closed once
	done      chan struct{} // no values sent; closed when Close is called
}

func (l *limitListener) release() {
	<-l.sem
	fmt.Println("release")
}
