package main

import "fmt"

type item struct {
	index, value int
}

func main() {

	integers := []int{2, 4, 6, 8, 10}

	ch := make(chan item, len(integers))

	for index, number := range integers {
		go pow2(item{index, number}, ch)
	}

	for range integers {
		newItem := <-ch
		integers[newItem.index] = newItem.value
	}

	fmt.Printf("RESULT: %v\n", integers)

}

func pow2(i item, ch chan item) {
	powed := i.value * i.value

	ch <- item{i.index, powed}
}
