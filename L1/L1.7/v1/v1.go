package v1

import (
	"fmt"
	"sync"
	"time"
)

// Ключ это название страны, значение - население в стране
// Реализация через Mutex
type World struct {
	mx sync.Mutex // !!!!!!!!!!!!!!!
	m  map[string]int
}

func (world *World) Load(key string) (int, bool) {
	world.mx.Lock()
	defer world.mx.Unlock()

	val, ok := world.m[key]

	return val, ok
}

func (world *World) Inc(key string) {
	world.mx.Lock()
	defer world.mx.Unlock()

	world.m[key]++
}

func NewWorld() *World {
	return &World{m: make(map[string]int)}
}

func workerRead(worl *World, country string) {
	for {
		time.Sleep(100 * time.Millisecond)
		people, ok := worl.Load(country)

		if !ok {
			fmt.Printf("No data at key: %v\n", country)
			continue
		}

		fmt.Printf("People at %v: %v\n", country, people)
	}
}

func workerWrite(world *World, country string) {
	for {
		time.Sleep(50 * time.Millisecond)
		world.Inc(country)
	}
}

func V1() {
	world := NewWorld()

	// Write

	go workerWrite(world, "RU")
	go workerWrite(world, "USA")
	go workerWrite(world, "CHINA")
	go workerWrite(world, "BRITANIA")
	go workerWrite(world, "UKRAIN")
	go workerWrite(world, "NATO")

	go workerWrite(world, "RU")
	go workerWrite(world, "USA")
	go workerWrite(world, "CHINA")
	go workerWrite(world, "BRITANIA")
	go workerWrite(world, "UKRAIN")
	go workerWrite(world, "NATO")

	go workerWrite(world, "RU")
	go workerWrite(world, "USA")
	go workerWrite(world, "CHINA")
	go workerWrite(world, "BRITANIA")
	go workerWrite(world, "UKRAIN")
	go workerWrite(world, "NATO")

	go workerWrite(world, "RU")
	go workerWrite(world, "USA")
	go workerWrite(world, "CHINA")
	go workerWrite(world, "BRITANIA")
	go workerWrite(world, "UKRAIN")
	go workerWrite(world, "NATO")

	go workerWrite(world, "RU")
	go workerWrite(world, "USA")
	go workerWrite(world, "CHINA")
	go workerWrite(world, "BRITANIA")
	go workerWrite(world, "UKRAIN")
	go workerWrite(world, "NATO")

	// Read
	go workerRead(world, "RU")
	go workerRead(world, "USA")
	go workerRead(world, "CHINA")
	go workerRead(world, "BRITANIA")
	go workerRead(world, "UKRAIN")
	go workerRead(world, "NATO")

	go workerRead(world, "RU")
	go workerRead(world, "USA")
	go workerRead(world, "CHINA")
	go workerRead(world, "BRITANIA")
	go workerRead(world, "UKRAIN")
	go workerRead(world, "NATO")

	timer := time.NewTimer(5 * time.Second)

	<-timer.C

}
