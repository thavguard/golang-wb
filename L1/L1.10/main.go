package main

import (
	"fmt"
	"sync"
)

type Group struct {
	mx sync.Mutex
	m  map[int][]float64
}

func (g *Group) Load(key int) ([]float64, bool) {
	g.mx.Lock()
	defer g.mx.Unlock()

	val, ok := g.m[key]

	return val, ok
}

func (g *Group) AddItemTo(key int, item float64) {
	g.mx.Lock()
	defer g.mx.Unlock()

	g.m[key] = append(g.m[key], item)

}

func NewGroup() *Group {
	return &Group{m: make(map[int][]float64)}
}

func main() {

	var wg sync.WaitGroup

	integers := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	group := NewGroup()

	for _, item := range integers {
		wg.Add(1)
		go worker(&wg, group, item)

	}

	wg.Wait()

	fmt.Printf("%v\n", group.m)

}

func worker(wg *sync.WaitGroup, g *Group, item float64) {
	defer wg.Done()

	base := (int(item) / 10) * 10
	g.AddItemTo(base, item)

}
