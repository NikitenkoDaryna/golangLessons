package main

import (
	"sync"
)

type Counter struct {
	v   map[string]Worker
	mux sync.RWMutex
}

func Start() Counter {
	return Counter{
		v: make(map[string]Worker),
	}
}
func (c *Counter) getCounter(id string) Worker {

	c.mux.RLock()

	defer c.mux.RUnlock()

	return c.v[id]

}

func (c *Counter) initCounter(id string,w Worker) {
	c.mux.Lock()

	c.v[id]=w
	c.mux.Unlock()

}
