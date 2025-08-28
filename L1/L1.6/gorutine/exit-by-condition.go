package gorutine

import (
	"fmt"
	"sync"
	"time"
)

func ExitByCondition() {

	var wg sync.WaitGroup

	wg.Add(1)
	go workerA(&wg)

	wg.Wait()

}

func workerA(wg *sync.WaitGroup) {
	defer wg.Done()

	counter := 0

	for {

		if counter > 5 {
			fmt.Printf("СТОП\n")
			break
		}

		time.Sleep(1 * time.Second)
		fmt.Printf("ГОРУТИНА РАБОТАЕТ...\n")

		counter++

	}
}
