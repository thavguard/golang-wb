package notificationsadapter

import (
	"L1.21/services/email"
	"L1.21/services/sms"
	"L1.21/services/telegram"
)

// SMS Adapter
type SmsAdapter struct {
	*sms.SmsService
	Target string
}

func (adapter *SmsAdapter) Send() bool {
	return adapter.SendSms(adapter.Target)
}

func NewSmsAdapter(smsSerivice *sms.SmsService, number string) Notifications {
	return &SmsAdapter{Target: number, SmsService: smsSerivice}
}

// Tg Adapter

type TelegramAdapter struct {
	*telegram.TelegramService
	Target string
}

func (adapter *TelegramAdapter) Send() bool {
	return adapter.SendMessage(adapter.Target)
}

func NewTelegramAdapter(tgSerivce *telegram.TelegramService, target string) Notifications {
	return &TelegramAdapter{Target: target, TelegramService: tgSerivce}
}

// Email Adapter

type EmailAdapter struct {
	*email.EmailService
	Email string
}

func (adapter *EmailAdapter) Send() bool {
	return adapter.SendEmail(adapter.Email)
}

func NewEmailAdapter(emailService *email.EmailService, email string) Notifications {
	return &EmailAdapter{Email: email, EmailService: emailService}
}
