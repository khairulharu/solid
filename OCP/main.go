package main

import (
	"errors"
)

type NotificationService interface {
	SendNotification(message string) (string, error)
}

type EmailNotificationService struct{}

func (e *EmailNotificationService) SendNotification(massage *string) (*string, error) {
	if massage != nil {
		return massage, nil
	}

	return nil, errors.New("no message")
}

func NewEmailNotificationService() *EmailNotificationService {
	return &EmailNotificationService{}
}
