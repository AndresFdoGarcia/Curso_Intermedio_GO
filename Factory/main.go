package main

import "fmt"

// SMS  *************************************************************
type INotificationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

// Implementar SendNotification para SMSNotification
func (SMSNotification) SendNotification() {
	fmt.Println("Sending notification via SMS")
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

// Implemetar GetSender para SMSNotification
func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

// Email *************************************************************
type EmailNotification struct {
}

// Implementar SendNotification para EmailNotification
func (EmailNotification) SendNotification() {
	fmt.Println("Sending notification via Email")
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

// Implemetar GetSender para EmailNotification
func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

// Factory *************************************************************
func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}

	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("No notification type")
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactiry, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotification(smsFactiry)
	sendNotification(emailFactory)

	getMethod(smsFactiry)
	getMethod(emailFactory)
}
