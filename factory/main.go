package main

import "fmt"

type INofiticationFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChannel() string
}

type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending notification via SMS")
}

func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChannel() string {
	return "Twilio"
}

type EmailNotification struct {
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

func (EmailNotification) SendNotification() {
	fmt.Println("Enviando notificacion vía Email")
}

// Isendere
type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChannel() string {
	return "SES"
}

//Factory
func getNotificationFactory(notificationType string) (INofiticationFactory, error) {
	if notificationType == "SMS" {
		return SMSNotification{}, nil
	}
	if notificationType == "Email" {
		return EmailNotification{}, nil
	}
	return nil, fmt.Errorf("No Notification Find")
}

func sendNotf(f INofiticationFactory) {
	f.SendNotification()
}

func getMethod(f INofiticationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	sendNotf(smsFactory)
	sendNotf(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)

}
