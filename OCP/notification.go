package main

type NotificationService interface {
	SendNotification(message *string) (*string, error)
}
