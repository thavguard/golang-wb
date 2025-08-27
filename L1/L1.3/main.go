package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(i int, ch chan int) {
	defer wg.Done()
	for item := range ch {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("GORUTINE: %v; VALUE: %v\n", i, item)
	}

}

func main() {
	// Канал для общения горутин
	ch := make(chan int)

	// Принимаем параментр workers из консоли
	workerCountParam := flag.String("workers", "4", "Число воркеров-горутин")
	flag.Parse()
	workerCount, err := strconv.Atoi(*workerCountParam)

	if err != nil {
		log.Fatal("-workers must be an integer")
	}

	// Запускаем N воркеров
	for i := range workerCount {
		wg.Add(1)
		go worker(i, ch)
	}

	// Продюсим бесконечный поток
	// В отдельной горутине чтобы не блокировать main горутину
	go func() {
		counter := 0
		for {
			time.Sleep(10 * time.Millisecond)
			counter++
			ch <- counter
		}

		close(ch) // В текущем кейсе бесконечной записи код не дойдет до `close`
		// но в других кейсах лучше закрывать

	}()

	// Wait чтобы не закрывалась программа
	wg.Wait()

}
