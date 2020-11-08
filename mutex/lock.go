package mutex

import "sync"

type Mutex struct {
	sync.Mutex
	safe bool
}

// safeLock := mutex.NewMutex(true)  safeLock.Lock() safeLock.Unlock()
func NewMutex(safe ...bool) *Mutex {
	mu := new(Mutex)
	if len(safe) > 0 {
		mu.safe = safe[0]
	} else {
		mu.safe = false
	}
	return mu
}

func (mu *Mutex) IsSafe() bool {
	return mu.safe
}

func (mu *Mutex) Lock() {
	if mu.safe {
		mu.Mutex.Lock()
	}
}

func (mu *Mutex) Unlock() {
	if mu.safe {
		mu.Mutex.Unlock()
	}
}
