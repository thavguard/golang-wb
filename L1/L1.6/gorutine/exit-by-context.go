package gorutine

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func ExitByContext() {

	var wg sync.WaitGroup

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT)
	defer cancel()

	wg.Add(1)
	go workerC(ctx, &wg)

	wg.Wait()

}

func workerC(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
ROOT:
	for {

		select {
		case <-ctx.Done():
			fmt.Printf("ПОЛУЧЕН СИГНАЛ ОБ ОСТАНОВКЕ\n")
			break ROOT

		default:
			time.Sleep(1 * time.Second)
			fmt.Printf("ГОРУТИНА РАБОТАЕТ...\n")
		}

	}
}
