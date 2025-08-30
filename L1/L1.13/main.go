package main

import "fmt"

func main() {
	swap1()
	swap2()
}

func swap1() {
	a := 10
	b := 5

	a = a + b
	b = a - b
	a = a - b

	fmt.Printf("A: %v; B: %v\n", a, b)
}

func swap2() {
	a := 10
	b := -5

	a = a ^ b
	b = b ^ a
	a = a ^ b

	fmt.Printf("A: %v; B: %v\n", a, b)
}
