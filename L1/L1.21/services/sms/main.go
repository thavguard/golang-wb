package sms

import (
	"fmt"
	"time"
)

type SmsService struct {
}

func (sms *SmsService) SendSms(number string) bool {
	time.Sleep(1 * time.Second)
	fmt.Printf("Сообщение улетело на номер %v\n", number)

	return true
}
