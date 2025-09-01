package notifications

import (
	"fmt"
	"sync"

	notificationsadapter "L1.21/adapters/notifications-adapter"
	"L1.21/services/email"
	"L1.21/services/sms"
	"L1.21/services/telegram"
)

type NotificationService struct {
}

func (n *NotificationService) SendToAll() {

	var wg sync.WaitGroup

	type MockUser struct {
		email string
		phone string
		tg    string
	}

	mockUser := &MockUser{email: "test@test.com", phone: "88005553535", tg: "durov"}

	smsAdapter := notificationsadapter.NewSmsAdapter(&sms.SmsService{}, mockUser.phone)
	tgAdapter := notificationsadapter.NewTelegramAdapter(&telegram.TelegramService{}, mockUser.tg)
	emailAdapter := notificationsadapter.NewEmailAdapter(&email.EmailService{}, mockUser.email)

	adapters := [3]notificationsadapter.Notifications{smsAdapter, tgAdapter, emailAdapter}

	for _, a := range adapters {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a.Send()

		}()
	}

	wg.Wait()
	fmt.Printf("Все уведомления отправлены!\n")

}
