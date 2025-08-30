package main

import (
	"fmt"

	myset "L.11/my-set"
)

func main() {
	myset1 := myset.NewMySet()
	myset2 := myset.NewMySet()

	myset1.Append(1, 2, 3)
	myset2.Append(2, 3, 4)

	intersect := myset1.Intersect(myset2.Set)

	fmt.Printf("Intersection: %v\n", intersect)
}
