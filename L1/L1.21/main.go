package main

import "L1.21/infra/notifications"

func main() {
	service := &notifications.NotificationService{}

	service.SendToAll()
}
