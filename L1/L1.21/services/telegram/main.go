package telegram

import (
	"fmt"
	"time"
)

type TelegramService struct {
}

func (t *TelegramService) SendMessage(target string) bool {
	time.Sleep(1 * time.Second)
	fmt.Printf("Сообщение улетело пользователю %v\n", target)

	return true
}
