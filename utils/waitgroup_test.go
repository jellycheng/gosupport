package utils

import (
	"fmt"
	"testing"
	"time"
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
	fmt.Println("end")
}

// go test -run=TestWaitGroupPoolV1
func TestWaitGroupPoolV1(t *testing.T) {
	wgp := NewWaitGroupPool(3)
	for i := 0; i < 60; i++ {
		wgp.Add()
		go func(x int) {
			defer wgp.Done()
			fmt.Println("hello-" + fmt.Sprintf("%d", x))
			time.Sleep(2 * time.Second)
		}(i)
	}

	wgp.Wait()
	fmt.Println("end")
}

// go test -run=TestWaitGroupPoolV2
func TestWaitGroupPoolV2(t *testing.T) {
	wgp := NewWaitGroupPoolV2(3)
	wgp.Wrap(func() {
		fmt.Println("任务1")
		time.Sleep(2 * time.Second)
	})
	for i := 0; i < 60; i++ {
		wgp.Wrap(func(x int) func() {
			return func() {
				fmt.Println("hello-" + fmt.Sprintf("%d", x))
				time.Sleep(2 * time.Second)
			}
		}(i))
	}

	wgp.Wait()
	fmt.Println("end")
}
