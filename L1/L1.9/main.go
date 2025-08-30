package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	array := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, 17711}

	chanX := make(chan int)
	chan2x := make(chan int)

	// Горутина пишущая элементы слайса в канал
	wg.Add(1)
	go func(sg *sync.WaitGroup, chx chan int, slice []int) {
		defer sg.Done()
		for _, fib := range slice {
			time.Sleep(100 * time.Millisecond)
			chx <- fib
		}

		close(chx)
	}(&wg, chanX, array)

	// Горутина прокси для x2 значений
	wg.Add(1)
	go func(sg *sync.WaitGroup, chx chan int, ch2x chan int) {
		defer sg.Done()
		for item := range chx {
			ch2x <- item * 2
		}

		close(ch2x)
	}(&wg, chanX, chan2x)

	// Пишушая в stdout горутина читающая из прокси
	wg.Add(1)
	go func(sg *sync.WaitGroup, ch2x chan int) {
		defer sg.Done()
		for item := range ch2x {
			fmt.Printf("2X: %v\n", item)
		}

	}(&wg, chan2x)

	wg.Wait()

}
