package gorutine

import (
	"fmt"
	"sync"
	"time"
)

func ExitByChannel() {
	var wg sync.WaitGroup
	exitChan := make(chan bool, 1)

	wg.Add(1)
	go workerE(&wg, exitChan)

	go func() {
		timer := time.NewTimer(5 * time.Second)

		<-timer.C
		exitChan <- true
	}()

	wg.Wait()

}

func workerE(wg *sync.WaitGroup, exitChan <-chan bool) {
	defer wg.Done()

ROOT:
	for {

		select {
		case <-exitChan:
			fmt.Printf("ПОЛУЧЕН СИГНАЛ ОБ ОСТАНОВКЕ!!\n")
			break ROOT

		default:
			time.Sleep(1 * time.Second)
			fmt.Printf("Горутина работает...\n")
		}

	}
}
