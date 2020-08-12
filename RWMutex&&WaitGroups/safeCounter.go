package main

import (
	"sync"
)

type Counter struct {
	v   map[Worker]int
	mux sync.RWMutex
}

func Start() Counter {
	return Counter{
		v: make(map[Worker]int),
	}
}
func (c *Counter) getCounter(w Worker) int {

	c.mux.RLock()

	defer c.mux.RUnlock()

	return c.v[w]

}

func (c *Counter) initCounter(wr Worker) {
	c.mux.Lock()

	c.v[wr] ++
	c.mux.Unlock()

}
