package email

import (
	"fmt"
	"time"
)

type EmailService struct {
}

func (s *EmailService) SendEmail(email string) bool {
	time.Sleep(1 * time.Second)
	fmt.Printf("Письмо ушло на почту %v\n", email)

	return true
}
