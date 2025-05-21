package utils

import (
	"sync"
)

type WaitGroupWrapper struct {
	sync.WaitGroup
}

func (w *WaitGroupWrapper) Wrap(callback func()) {
	w.Add(1)
	go func() {
		defer w.Done()
		callback()
	}()
}

func NewWaitGroup() WaitGroupWrapper {
	return WaitGroupWrapper{}
}
