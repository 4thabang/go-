package main

import (
	"sync"
)

type Counter struct {
	mu  sync.Mutex
	num int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.num++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	return c.num
}
