package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ch chan int) {
	defer wg.Done()

	for item := range ch {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("GOOD JOB: %v\n", item)
	}

}

func main() {

	ch := make(chan int, 64) // Создаем канал с длинной буфера 64 для примера того, что после завершения работы программа безопастно дожидается завершения всех горутин
	counter := 0

	wg.Add(2)
	go worker(ch)
	go worker(ch)

	timer := time.NewTimer(3 * time.Second)

ROOT:
	for {

		select {
		case <-timer.C:
			fmt.Printf("TIME TO LEAVE\n")
			break ROOT

		default:
			time.Sleep(10 * time.Millisecond)
			counter++
			ch <- counter

		}

	}

	fmt.Printf("CURRENT COUNTER IS: %v\n", counter)
	close(ch)

	wg.Wait()
}
