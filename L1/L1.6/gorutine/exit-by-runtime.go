package gorutine

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func ExitByRuntime() {
	var wg sync.WaitGroup

	wg.Add(1)
	go runtimeWorker(&wg)

	wg.Wait()
	fmt.Printf("Программа завершилась...\n")
}

func runtimeWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	counter := 0
	for {

		if counter >= 5 {
			fmt.Printf("Выход из горутины\n")
			runtime.Goexit()
		}

		time.Sleep(1 * time.Second)
		fmt.Printf("Горутина работает...\n")

		counter++
	}
}
