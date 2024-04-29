package main

import "fmt"

type CartNotificationService struct{}

func (c *CartNotificationService) SendNotification(message *string) (*string, error) {
	fmt.Printf("Cart Notifcation: %s \n", *message)
	return nil, nil
}
