package utils

import (
	"fmt"
	"testing"
)

// go test -run=TestWaitGroupWrapper_Wrap
func TestWaitGroupWrapper_Wrap(t *testing.T) {
	wg := NewWaitGroup()
	// wg := WaitGroupWrapper{}
	// wg := &WaitGroupWrapper{}

	wg.Wrap(func() {
		fmt.Println("hello 01")
	})

	wg.Wrap(func() {
		fmt.Println("hello 02")
	})
	wg.Wrap(func() {
		fmt.Println("hello 03")
	})

	wg.Wait()
}
