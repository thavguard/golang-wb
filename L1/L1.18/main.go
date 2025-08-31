package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mx sync.Mutex
	c  int
}

func (c *Counter) Inc() {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.c++
}

func (c *Counter) Load() int {
	c.mx.Unlock()
	defer c.mx.Unlock()

	return c.c

}

func newCounter() *Counter {
	return &Counter{}
}

func main() {
	var wg sync.WaitGroup

	counter := newCounter()

	for range 50 {
		wg.Add(1)
		go worker(&wg, counter)
	}

	wg.Wait()

	fmt.Printf("COUNTER: %v\n", counter.c)
}

func worker(wg *sync.WaitGroup, c *Counter) {
	defer wg.Done()

	for range 5 {
		c.Inc()
	}

}
