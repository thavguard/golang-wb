package v2

import (
	"fmt"
	"sync"
	"time"
)

// Ключ это название страны, значение - население в стране
// Реализация через sync.Map
type World struct {
	sync.Map // !!!!!!!!!!
}

func (world *World) Inc(key string) {

	var prevN int

	if prev, ok := world.Load(key); ok {
		prevN = prev.(int)

		world.Store(key, prevN+1)
	} else {
		world.Store(key, 1)
	}

}

func NewWorld() *World {
	return &World{}
}

func workerRead(worl *World, country string) {
	for {
		time.Sleep(100 * time.Millisecond)

		var people int

		if _people, ok := worl.Load(country); ok {
			people = _people.(int)
		} else {
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

func V2() {
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
