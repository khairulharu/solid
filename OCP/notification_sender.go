package main

type NotificationSender struct {
	notificationService NotificationService
}

func (n *NotificationSender) SendNotification(message *string) (*string, error) {
	return n.notificationService.SendNotification(message)
}
