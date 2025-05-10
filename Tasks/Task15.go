package main

import "fmt"

type Notifier interface {
	Notify()
}

type EmailNotifier struct {
	emain string
}

type SMSNotifier struct {
	sms string
}

func (e EmailNotifier) Notify() {
	fmt.Println("Emain Notify")
}

func (s SMSNotifier) Notify() {
	fmt.Println("SMS Notify")
}

func main() {
	email := EmailNotifier{}
	sms := SMSNotifier{}
	email.Notify()
	sms.Notify()
}
