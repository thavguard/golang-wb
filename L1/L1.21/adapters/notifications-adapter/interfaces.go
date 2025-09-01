package notificationsadapter

// Интерфейс работы с уведомлениями
type Notifications interface {
	Send() bool
}
