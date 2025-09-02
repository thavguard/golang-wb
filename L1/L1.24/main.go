package main

import (
	"fmt"

	"L1.24/coordinates"
)

func main() {
	p1 := coordinates.NewPoint(10, 5)
	p2 := coordinates.NewPoint(5, 10)

	fromP1ToP2 := p1.Distance(p2)

	fmt.Printf("fromP1ToP2: %v\n", fromP1ToP2)
}
