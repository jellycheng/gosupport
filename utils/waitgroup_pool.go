package utils

import (
	"math"
	"sync"
)

type WaitGroupPool struct {
	pool chan struct{}
	wg   *sync.WaitGroup
}

func NewWaitGroupPool(size int) *WaitGroupPool {
	if size <= 0 {
		size = math.MaxInt32
	}
	return &WaitGroupPool{
		pool: make(chan struct{}, size),
		wg:   &sync.WaitGroup{},
	}
}

func (p *WaitGroupPool) Add() {
	p.pool <- struct{}{}
	p.wg.Add(1)
}

func (p *WaitGroupPool) Done() {
	<-p.pool
	p.wg.Done()
}

func (p *WaitGroupPool) Wait() {
	p.wg.Wait()
}

type WaitGroupPoolV2 struct {
	pool chan struct{}
	wg   *sync.WaitGroup
}

func NewWaitGroupPoolV2(size int) *WaitGroupPoolV2 {
	if size <= 0 {
		size = math.MaxInt32
	}
	return &WaitGroupPoolV2{
		pool: make(chan struct{}, size),
		wg:   &sync.WaitGroup{},
	}
}

func (p *WaitGroupPoolV2) Wrap(callback func()) {
	p.pool <- struct{}{}
	p.wg.Add(1)
	go func() {
		defer func() {
			<-p.pool
			p.wg.Done()
		}()
		callback()
	}()
}

func (p *WaitGroupPoolV2) Wait() {
	p.wg.Wait()
}
