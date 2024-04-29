package main

import "fmt"

type EmailNotificationService struct{}

func (e *EmailNotificationService) SendNotification(message *string) (*string, error) {

	fmt.Printf("Email Notifcation: %s \n", *message)
	return nil, nil
}
