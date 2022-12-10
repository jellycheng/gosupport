package mutex

import (
	"fmt"
	"testing"
	"time"
)

func TestNewMutex(t *testing.T) {
	c := make(map[string]int)
	safeLock := NewMutex(true)
	go func() { //开一个goroutine写map

		for j := 0; j < 1000000; j++ {
			safeLock.Lock()
			c[fmt.Sprintf("%d", j)] = j
			safeLock.Unlock()
		}
	}()

	go func() { //再开一个goroutine读map
		for j := 0; j < 1000000; j++ {
			safeLock.Lock()
			fmt.Println(c[fmt.Sprintf("%d", j)])
			safeLock.Unlock()
		}
	}()

	time.Sleep(time.Second * 20)

}
