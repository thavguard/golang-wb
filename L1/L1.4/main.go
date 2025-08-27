package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context, i int, ch chan int) {
	defer wg.Done()

ROOT:
	for item := range ch {
		time.Sleep(1 * time.Second)
		fmt.Printf("GORUTINE: %v; VALUE: %v\n", i, item)

		select {
		case <-ctx.Done():
			fmt.Printf("ПОЛУЧЕН СИГНАЛ НО В КАНАЛЕ ВСЕ ЕЩЕ ЕСТЬ СООБЩЕНИЯ ПОЭТОМУ ОБРАБОТАЕМ ИХ ПЕРЕД ВЫХОДОМ\n")

		default:
			continue ROOT
		}

	}

}

// Комментарии по работе самой программы доступны в файле “L1/L1.3/main.go“.
// Тут комментарии только про сигналы
func main() {

	// Создаем контекст отмены и слушаем два сигнала
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT)
	defer stop() // откладываем завершение программы до момента пока main не исполнится

	ch := make(chan int)

	workerCountParam := flag.String("workers", "69", "Число воркеров-горутин")
	flag.Parse()
	workerCount, err := strconv.Atoi(*workerCountParam)

	if err != nil {
		log.Fatal("-workers must be an integer")
	}

	for i := range workerCount {
		wg.Add(1)
		go worker(ctx, i, ch)
	}

	go func() {
		counter := 0

		// В цикле пушим сообщения в канал, если ловим завершение останавливаем продьюсинг и закрываем канал
		// При этом горутины продолжают читать те данные которые там остались и корректно завершают свою работу
		// благодаря sync.WaitGroup
	ROOT:
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("STOP PROGRAM\n")
				break ROOT

			case <-time.After(100 * time.Millisecond):
				counter++
				ch <- counter
			}

		}

		close(ch) // Если не закрыть канал горутины не узнают о том что большие сообщения мы не продьюсим

	}()

	wg.Wait() // Ждем пока все горутины отработают

	fmt.Printf("Программа корректно завершила все операции. Bye bye...\n")

}
