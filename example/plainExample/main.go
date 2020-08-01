package main

import (
	"fmt"
	"gomailer"
	"gomailer/pkg/models"
)

var smtpMail, smtpPassword, smtpHost = "email@gmail.com", "emailpassword", "smtp.gmail.com"

func main() {
	config := models.Configuration{SmtpMail: smtpMail, SmtpPassword: smtpPassword, SmtpHost: smtpHost}
	mailer := gomailer.New(config)
	mail := models.Email{To: []string{"email@gmail.com", smtpMail}, Subject: "sup bro", Body: "im basic email", From: "email@gmail.com"}
	sent, validationErr, err := mailer.SendPlainEmail(mail)
	fmt.Println("result: %v,%v,%v", sent, validationErr, err)
}
